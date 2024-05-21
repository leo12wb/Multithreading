# Desafio de Multithreading e APIs em Go

Este é um projeto de exemplo para demonstrar como fazer solicitações a duas APIs diferentes em Go usando goroutines para aproveitar o paralelismo.

## Requisitos

Certifique-se de ter Go instalado em seu sistema. Você pode encontrar as instruções de instalação em [golang.org](https://golang.org/doc/install).

## Como executar

1. Clone o repositório:

## Detalhes do Desafio

O desafio consiste em fazer solicitações simultâneas a duas APIs diferentes para obter dados de endereço com base em um CEP fornecido. O objetivo é obter a resposta mais rápida entre as duas APIs e descartar a resposta mais lenta. O tempo de resposta é limitado a 1 segundo. Se nenhuma resposta for recebida dentro desse tempo limite, um erro de timeout será exibido.

As APIs usadas neste projeto são:

- [Brasil API](https://brasilapi.com.br/api/cep/v1/) - para obter dados de endereço com base no CEP.
- [ViaCEP](http://viacep.com.br/ws/) - também para obter dados de endereço com base no CEP.

## Contribuindo

Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para obter mais detalhes.
