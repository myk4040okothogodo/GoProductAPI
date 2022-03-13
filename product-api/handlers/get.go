package handlers

import (
    "context"
    "net/http"
    "github.com/myk4040okothogodo/Grpc2/product-api/data"
    protos "github.com/myk4040okothogodo/Grpc2/currency/protos/currency/protos"
)


//ListALl handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request){
    p.l.Println("[DEBUG] get all records")
    rw.Header().Add("Content-Type", "application/json")
    prods := data.GetProducts()

    err := data.ToJSON(prods, rw)
    if err != nil {
        //we should never be here but log the error just incase
        p.l.Println("[ERROR] serializing product", err)
    }
}

//ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request){
    rw.Header().Add("Content-Type", "application/json")
    id := getProductID(r)

    p.l.Println("[DEBUG] get record id", id)
    prod, err := data.GetProductByID(id)

    switch err {
    case nil:

    case data.ErrProductNotFound:
        p.l.Println("[ERROR] fetching product", err)
        rw.WriteHeader(http.StatusNotFound)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return
    default:
        p.l.Println("[ERROR] fetching product", err)
        rw.WriteHeader(http.StatusInternalServerError)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)
        return


    }
    //get exchange rate
    rr := &protos.RateRequest{
        Base: protos.Currencies(protos.Currencies_value["EUR"]),
        Destination: protos.Currencies(protos.Currencies_value["GBP"]),

    }
    resp, err := p.cc.GetRate(context.Background(), rr)
    if err != nil {
        p.l.Println("[Error] error getting new rate", err)
        data.ToJSON(&GenericError{Message: err.Error()}, rw)

    }
    p.l.Printf("Resp %#v", resp)
    prod.Price = prod.Price * resp.Rate

    err = data.ToJSON(prod, rw)
    if err != nil {
        //we should never be here but log the error just incase
        p.l.Println("[ERROR] serializing product", err)

    }

}



