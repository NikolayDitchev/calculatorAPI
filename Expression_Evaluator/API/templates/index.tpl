<!DOCTYPE html>
<html>
    <head>
    </head>
    <body>
        <input type="text" name="expression" id="expression">
        <input type="button" value="submit" onclick="submit()" name="submit" id="submit">    
        <p id="answer"></p>
        <script >
            function submit(){

                expressionString = document.getElementById('expression').value

                fetch("http://localhost:8080/evaluate", {
                    method: "POST",
                    body: JSON.stringify({
                        expression: expressionString
                    }),
                    headers: {
                        "Content-Type": "application/json"
                    }
                })
                .then((response) => response.text())
                .then((text) => {
                    document.getElementById('answer').innerHTML = text
                })
            }
        </script>
    </body>
</html>