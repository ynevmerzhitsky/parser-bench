package main

import (
    "fmt"
    "strings"
    "github.com/oleksandr/conditions"
    "github.com/nullne/evaluator"
    "github.com/Knetic/govaluate"
    "testing"
)

var (

    TestFunc = conditionsTest
    data = []string{"1", "2", "3", "4", "5", "10"}
    deRecord = map[string]interface{}{
        "product:version": 1,
        "version":         1,
        "device:gpu":      "Intel Iris,GeForce",
        "$1":              "Intel Iris,GeForce",
        "$0":              1,
        "arg0":            1,
    }
    result interface{}
)

func evaluatorTest(record map[string]interface{}) (error) {
    params := evaluator.MapParams(record)
    in := strings.Join(data, " ")
    var testErr error
    result, testErr = evaluator.EvalBool(
        fmt.Sprintf(`(and (not (in product:version (%s))) (in product:version (%s)) (in product:version (%s)))`, in, in, in), params)
    if testErr != nil {
        return testErr
    }
    return nil
}

func govaluateTest(record map[string]interface{}) (error) {
    in := strings.Join(data, ",")
    expression, testErr := govaluate.NewEvaluableExpression(fmt.Sprintf("!(arg0 IN (%s)) && (arg0 IN (%s)) && (arg0 IN (%s))", in, in, in))
    if testErr != nil {
        return testErr
    }
    result, testErr = expression.Evaluate(deRecord)
    return testErr
}

func conditionsTest(record map[string]interface{}) (error) {
    in := strings.Join(data, ",")
    expressionStr := fmt.Sprintf("($0 NOT IN [%s]) AND ($0 IN [%s]) AND ($0 IN [%s])", in, in, in)
    p := conditions.NewParser(strings.NewReader(expressionStr))
    expr, testErr := p.Parse()
    if testErr != nil {
        return testErr
    }
    result, testErr = conditions.Evaluate(expr, deRecord)
    if testErr != nil {
        return testErr
    }
    return nil
}

func BenchmarkExpression(b *testing.B) {
    for i:=0;i<b.N; i++ {
        err := TestFunc(deRecord)
        if err != nil {
            b.Error(err)
        }

    }
}
