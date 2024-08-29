package handlers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

// Testa a função HelloHandler para garantir que ela responde corretamente
func TestHelloHandler(t *testing.T) {
    // Cria uma requisição HTTP de teste do tipo GET
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Cria um ResponseRecorder para capturar a resposta do handler
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HelloHandler) // Converte a função HelloHandler para um HandlerFunc

    // Chama o handler com a requisição de teste
    handler.ServeHTTP(rr, req)

    // Verifica se o status da resposta é 200 OK
    assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

    // Verifica se o corpo da resposta é o esperado
    assert.Equal(t, "Hello from the handler!", rr.Body.String(), "Expected response body")
}
