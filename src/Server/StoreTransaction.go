//  StoreTransaction.go  -  Update/query App store transactions
//
//  E.B.Smith  -  December, 2015.


package main


import (
    "time"
    "errors"
    "violent.blue/GoKit/Log"
    "violent.blue/GoKit/Util"
    "violent.blue/GoKit/pgsql"
)


type StoreTransaction struct {
    TransactionID       string
    StoreID             string
    StoreTransactionID  string
    UserID              string
    Quantity            int
    Purchase            string
    PurchaseDate        time.Time
    Locale              string
    LocalizedPrice      string
}


func StoreTransactionWithTransactionID(storeID string, transactionID string) (*StoreTransaction, error) {
    Log.LogFunctionName()

    row := config.DB.QueryRow(
        `select
             transactionID
            ,storeID
            ,storeTransactionID
            ,userID
            ,quantity
            ,purchase
            ,purchaseDate
            ,locale
            ,localizedPrice
            from StoreTransactionTable where storeID = $1 and storeTransactionID = $2;`,
            storeID, transactionID)

    var st StoreTransaction
    error := row.Scan(
        st.TransactionID,
        st.StoreID,
        st.StoreTransactionID,
        st.UserID,
        st.Quantity,
        st.Purchase,
        st.PurchaseDate,
        st.Locale,
        st.LocalizedPrice)
    if error != nil {
        Log.LogError(error)
        return nil, error
    }

    return &st, nil
}


func StoreTransactionInsert(transaction StoreTransaction) (*StoreTransaction, error) {
    Log.LogFunctionName()

    transaction.TransactionID = Util.NewUUIDString()
    result, error := config.DB.Exec(
        `insert into StoreTransactionTable
            (transactionID
            ,storeID
            ,storeTransactionID
            ,userID
            ,quantity
            ,purchase
            ,purchaseDate
            ,locale
            ,localizedPrice)
            values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
        transaction.TransactionID,
        transaction.StoreID,
        transaction.StoreTransactionID,
        transaction.UserID,
        transaction.Quantity,
        transaction.Purchase,
        transaction.PurchaseDate,
        transaction.Locale,
        transaction.LocalizedPrice)
    if pgsql.RowsUpdated(result) != 1 && error == nil {
        error = errors.New("No rows updated")
    }
    if error != nil {
        Log.LogError(error)
        return nil, error
    }
    return &transaction, nil
}

