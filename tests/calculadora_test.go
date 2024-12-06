package soma

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

var ResultadoOperacao int

func calculaValores(v1 int, operador string, v2 int) error {
	switch operador {
	case "+":
		ResultadoOperacao = v1 + v2
	case "-":
		ResultadoOperacao = v1 - v2
	case "/":
		if v2 == 0 {
			return errors.New("divisão por zero não é permitida")
		}
		ResultadoOperacao = v1 / v2
	case "*":
		ResultadoOperacao = v1 * v2
	default:
		return errors.New("operador inválido")
	}

	return nil
}

func oValorDeveSer(resultadoEsperado int) error {
	if ResultadoOperacao != resultadoEsperado {
		return fmt.Errorf("resultado esperado %d, mas obtido %d", resultadoEsperado, ResultadoOperacao)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^realizo a operação (\d+) ([\+\-\*/]) (\d+)$`, calculaValores)
	ctx.Step(`^Então o valor deve ser (\d+)$`, oValorDeveSer)
}

// Integração com `testing.T`
func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Paths:    []string{"../features"}, // Caminho para o arquivo .feature
			Format:   "pretty",                // Formato do relatório de saída
			Strict:   true,
			TestingT: t,
		},
	}
	// test := &testUtils{
	// 	uri:    "https://httpbin.org/get",
	// 	client: &http.Client{},
	// }

	if suite.Run() != 0 {
		t.Fail() // Indica falha se o suite retornar erro
	}
}
