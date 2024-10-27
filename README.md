# Go Weather App - OTEL

Um serviço de previsão do tempo que recebe um CEP para buscar a localização e retorna a temperatura em Celsius, Kelvin e Fahrenheit.

Utiliza as APIs do [ViaCEP](https://viacep.com.br/) e [WeatherAPI](weatherapi.com/).

## Serviços

O projeto é composto por 2 serviços:

- **Serviço A:** Responsável por receber um CEP, fazer as devidas verificiações e chamar o Serviço B.

- **Serviço B:** Responsável por receber a cidade e retornar a temperatura.

- Exemplo de Chamada feita para o serviço A:
  
```bash
curl -X POST http://localhost:8000/15055285

```

- Exemplo de Resposta final após o processamento feito pelo serviço B:

```json
{
    "city": "São José do Rio Preto",
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

## Zipkin

Para visualizar as chamadas entre os serviços por meio dos spans, acesse o Zipkin em:

```bash
http://localhost:9411/zipkin/
```
