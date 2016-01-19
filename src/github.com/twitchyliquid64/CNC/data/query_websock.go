package data

import (
  "github.com/twitchyliquid64/CNC/logging"
  "github.com/twitchyliquid64/CNC/config"
  "golang.org/x/net/websocket"
  "database/sql"
	_ "github.com/lib/pq"
  "reflect"
)

func SqlQueryServer(ws *websocket.Conn){

  dbConn, err := sql.Open("postgres", "postgres://" + config.All().Database.Username +
                                     ":" + config.All().Database.Password +
                                     "@" + config.All().Database.Address +
                                     "/" + config.All().Database.Name +
                                     "?sslmode=require")
  if err != nil{
    logging.Error("data-websock", "DB error: ", err.Error())
    return
  }
  defer dbConn.Close()

  for {
    var data struct {
      Type string
      Query string
    }
    err := websocket.JSON.Receive(ws, &data)
    if err != nil{
      logging.Warning("data-websock", "Recieve error: ", err.Error())
      return
    }

    //process message
    switch data.Type {
      case "query":
        logging.Info("data-websock", "Got Query: ", data.Query)
        doQuery(data.Query, ws, dbConn)
    }
  }
}

func doQuery(sql string, ws *websocket.Conn, db *sql.DB) {
  rows, err := db.Query(sql)
  if err != nil {
    sendError(ws, err)
    return
  }
  defer rows.Close()

  cols, err := rows.Columns()
  if err != nil {
    sendError(ws, err)
    return
  }

  //send cols
  websocket.JSON.Send(ws, struct{
    Type string
    Cols []string
  }{Type: "columns", Cols: cols,})

  var resultSet [][]interface{}

  for rows.Next() {
    //create slice to store results for the row in
    var results []interface{}
    for i := 0; i < len(cols); i++ {
      var n interface{}
      results = append(results, &n)
    }

    err = rows.Scan(results...)
    if err != nil {
      sendError(ws, err)
      return
    }

    for i := 0; i < len(results); i++{
      if reflect.TypeOf(*(results[i].(*interface{}))) == reflect.TypeOf([]uint8{}){
        results[i] = string((*results[i].(*interface{})).([]uint8))
      }
    }

    resultSet = append(resultSet, results)
  }

  //send result set
  websocket.JSON.Send(ws, struct{
    Type string
    Results [][]interface{}
  }{Type: "data", Results: resultSet,})

}

func sendError(ws *websocket.Conn, e error){
  websocket.JSON.Send(ws, struct{
    Type string
    Error string
  }{Type: "error", Error: e.Error(),})
}
