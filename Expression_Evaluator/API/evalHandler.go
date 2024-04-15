package main

import (
	"Expression_Evaluator/internal/eval"
	"encoding/json"
	"net/http"
	"strconv"
)

// add fields to this struct if more parameters have to be accepted by the request via JSON
type exprStruct struct {
	Expression string
}

func evaluateExpression(writer http.ResponseWriter, req *http.Request) {

	var exprInput exprStruct

	jsonDecoder := json.NewDecoder(req.Body)
	errJson := jsonDecoder.Decode(&exprInput)

	if errJson != nil {
		http.Error(writer, errJson.Error(), http.StatusBadRequest)
		return
	}

	if exprInput.Expression == "" {
		http.Error(writer, "expression value missing", http.StatusBadRequest)
		return
	}

	if jsonDecoder.More() {
		http.Error(writer, "unneeded fields are provided", http.StatusBadRequest)
	}

	result, errEvaluate := eval.Evaluate(exprInput.Expression)

	if errEvaluate != nil {
		writer.Write([]byte(errEvaluate.Error()))
		return
	}

	var resultString string = strconv.FormatFloat(result, 'f', 3, 64)

	writer.Write([]byte(resultString))
}
