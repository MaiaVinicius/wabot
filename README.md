# Wabot

## Introdução

O objetivo desse projeto é enviar e receber mensagens utilizando o WhatsApp


## Como utilizar

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

## Recomendações

- Utilize o WhatsApp Business.
- Não tente rodar esse App com fins de Marketing - Acredite, o WhatsApp irá te bloquear muito rápido.
- Para cada mensagem, tente variar o conteúdo. Dessa forma ajuda o WhatsApp a entender como um mensagem mais "humana".
- Mantenha o celular que contém o CHIP **sempre** conectado no wi-fi.
- Preferencialmente, mantenha a tela do celular sempre ligada, conectado a um carregador e com o WhatsApp Business sempre aberto.


## Roadmap de novas funções

- Asegurar que a aplicação fique rodando 24/7.
- Criação de um Dashboard para controlar os envios / respostas
- Hospedar em um server
- Adaptar para rodar mais de uma instância com CHIP de disparo diferente - para aumentar a taxa de envio
- Tentar encontrar um timeout menor que não seja bloqueado.
- Desenvolver envio de mensagem baseado na resposta do usuário - ChatBot

## Comparação com a api do WhatsApp Business

| Plataforma | Valor / envio | Taxa de habilitação | Risco de bloqueio do número | Permite Marketing | Modelo fixo de mensagem | Disparo / min. | Captura de resposta |  Envio deme imag | Permite envio de mais um número |
|:-----------------------------:|---------------|---------------------|----------------------------------|--------------------------------|-------------------------------|----------------|------------------------|------------------|-----------------------------------|
| WhatsApp Business API oficial | 0,36 | 6.000,00 | Não | Não | Sim | Ilimitado | Sim - por WebHook | Sim | Não - limitado a apenas um número |
| WABOT | 0,00 | 17,00 | Sim, se utilizar de forma errada | Sim - mas poderá ser bloqueado | Não - permite qualquer modelo | 2 | Sim - a cada 5 minutos | Não | Quantos números precisar |