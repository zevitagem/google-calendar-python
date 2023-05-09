# google-calendar-python

## Documentação
- GoLang [https://developers.google.com/calendar/api/quickstart/go?hl=pt-br]
- Python [https://developers.google.com/calendar/api/quickstart/python?hl=pt-br]

## Comandos
```
python core.py list
python core.py insert
python core.py update <event_id>
```

## Fluxo
- Se encontrou o arquivo `token.json` e for válido, utilize-o
- Se não encontrou o arquivo `token.json`, ou encontrou, mas não é válido:
    - Caso encontre o arquivo, possui a data de expiração vencida e possui o `refresh token`, então refaça o token.
    - Qualquer negativa da condição anterior, redirecione o usuário para a tela de login e autenticação do Google.
- Caso seja gerado um novo token, seja por `refresh` ou pelo redirecionamento em tempo de execução, será feita a escrita de um novo arquivo `token.json` para que o processo não precise ser refeito toda vez que executado.

## Exemplo do arquivo de `credentials.json`:

```json
{
    "installed": {
        "client_id": "<code>.apps.googleusercontent.com",
        "project_id": "<number>",
        "auth_uri": "https://accounts.google.com/o/oauth2/auth",
        "token_uri": "https://oauth2.googleapis.com/token",
        "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
        "client_secret": "<secret>",
        "redirect_uris": [
            "http://localhost"
        ]
    }
}
```

## Exemplo do arquivo de `token.json`:

```json
{
    "token": "<token>",
    "refresh_token": "1//<refresh_token>",
    "token_uri": "https://oauth2.googleapis.com/token",
    "client_id": "<code>.apps.googleusercontent.com",
    "client_secret": "<secret>",
    "scopes": [
        "https://www.googleapis.com/auth/calendar"
    ],
    "expiry": "2023-05-08T13:22:55.592259Z"
}
```
