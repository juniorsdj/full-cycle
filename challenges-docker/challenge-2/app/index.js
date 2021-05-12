const app = require("express")();
const mysql = require("mysql2");

const connection = mysql.createConnection({
  host: "mysqlcontainer",
  user: "root",
  password: "pass",
  database: "pfa",
});

app.get("/:username", (req, resp) => {
  const { username } = req.params;
  console.log(username);

  connection.query(
    `INSERT INTO people SET nome=?`, username,
    (err, results) => {
      if (err) {
        return resp.send(`Deu ruim :(  
          ${err.message}`);
      }
      connection.query(
        "SELECT * FROM `people`",
        function (err, results, fields) {
          if (err) {
            return resp.send(`Deu ruim :(  
            ${err.message}`);
          }
          return resp.send(`
        <h1>Full Cycle Rocks!</h1>
        <ul>
        ${results.reduce((agg, cur) => agg + `<li>${cur.nome}</li>`, "")}
        </ul>`);
        }
      );
    }
  );
});

const port = 3000;

app.listen(port, () => {
  console.log(`app is running on ${port}`);
});
