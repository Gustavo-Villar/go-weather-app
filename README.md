# Go Weather App

Um serviço de previsão do tempo que recebe um CEP para buscar a localização e retorna a temperatura em Celsius, Kelvin e Fahrenheit.

Utiliza as APIs do [ViaCEP](https://viacep.com.br/) e [WeatherAPI](weatherapi.com/).

## Modo de uso

Para utilizar o serviço, basta enviar uma requisição GET para o endpoint `/weather` com o CEP desejado.

- Exemplo de Chamada:
  
```bash
curl -X GET "http://localhost:8080/weather?cep=15055285"

```

- Exemplo de Resposta:

```json
{
    "temp_C": 25.8,
    "temp_F": 78.44,
    "temp_K": 298.95
}
```

## Executando o projeto

Para executar o projeto use o comando para executar os testes e subir a api:

```bash
docker-compose up --build
```

## Google Run URL

A aplicação está disponível no Google Run, o link base da aplicação é <https://myapp-868620356628.us-central1.run.app> e as chamadas podem ser feitas da mesma forma que localmente.

Exemplo de chamada:

```bash
curl -X GET "https://myapp-868620356628.us-central1.run.app/weather?cep=15055285"

```
