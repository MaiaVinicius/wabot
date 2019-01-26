# Wabot
O objetivo desse projeto é enviar e receber mensagens utilizando o WhatsApp

Como utilizar:

- Renomeie o arquivo `.env.example` para `.env` na pasta raiz do projeto
- Preencha as informações de conexão com o banco de dados
- Preencha os endpoints:
    - Que fornecerá a fila de envio (`QUEUE_URL`)
    - Que removerá um envio da fila no servidor (`REMOVE_QUEUE_URL`)
    - Que receberá as respostas (`RESPONSES_URL`)
- Preencha as tabelas `wabot_project` e `wabot_sender`
- Rode o arquivo main.go 
    - Obs: Da primeira vez que rodar a aplicação irá exibir um QR Code na tela. Escaneie esse QR Code com seu WhatsApp
- A aplicação irá receber as mensagens e depois enviar cada mensagem da fila em cerca de 30-50 segundos cada mensagem

# Recomendações

- Utilize o WhatsApp Business.
- Não tente rodar esse App com fins de Marketing - Acredite, o WhatsApp irá te bloquear muito rápido.
- Para cada mensagem, tente variar o conteúdo. Dessa forma ajuda o WhatsApp a entender como um mensagem mais "humana".

