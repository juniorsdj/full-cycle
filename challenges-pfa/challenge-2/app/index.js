const app = require("express")();
const mysql = require('mysql2');


const connection = mysql.createConnection({
  host: 'mysqlcontainer',
  user: 'root',
  password: "pass",
  database: 'pfa'
});
 

app.get("/", (req, resp) => {
  connection.query(
    'SELECT * FROM `module`',
    function(err, results, fields) {
      if(err){
        return resp.send(`Deu ruim :(  
          ${err.message}`)
      }
      return resp.send(results);
    }
  );
  
});



const port = 3000;

app.listen(port, () => {
  console.log(`app is running on ${port}`);
});
