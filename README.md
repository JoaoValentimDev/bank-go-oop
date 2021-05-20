# Orientação a Objetos em Golang

Esse projeto tem como principal função, demonstrar a orientação a objetos na linguagem Go. Algo de extrema importancia
e que se encontra presente em diversas aplicações atualmente.
A oritentação a objetos alem de ser algo que aproxima nossos projetos do mundo real, da uma base solida na construção
dos mesmos.

## As "class" em Go

Pois é. Em Go existem as **structs**, estruturas para os tipos que construimos. Não existe a palavra reservada ***"class"***.
Veja um exemplo do proprio programa que desenvolvi:

```go

type Holder struct {
	Name, CPF, Profession string
}

```

Neste exemplo criamos um tipo Nomeado "Holder", que é uma estrutura composta por campos do tipo string.

## Como usar as structs

Declaramos da seguinte maneira:

```go
holder := clients.Holder{
	Name:       "João",
	CPF:        "222.222.222-22",
	Profession: "Developer",
}

currentAccount := accounts.CurrentAccount{
	Holder:        holder,
	AgencyNumber:  2,
	AccountNumber: 2,
}

fmt.Println(currentAccount, holder)
```