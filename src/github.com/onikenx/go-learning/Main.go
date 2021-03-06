package main

import (
	//biblioteca para controlar strings
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	// "io/ioutil"
	// "log"
	// "math"
	// "sort"
	// "strcov"
)

////////////////////////////////////////////////////////////

//variaveis que estejam em maiúscula no inicio são visíveis
//fora do pacote

var I int = 26

func variables() {
	// podemos declarar variaveis destas seguintes formas
	var j int
	i := 42
	j = 27
	fmt.Println(i + j)
	fmt.Println("Problem?")
	//%v printa a variavel e %T o tipo
	fmt.Printf("%v, %T", j, j)
}

////////////////////////////////////////////////////////////

func convertingVar() {
	var i float32 = 42.5
	var j int = 5
	//funca
	j = int(i)
	//nao funca porque irias estar a perder info
	//j = i
	//devia funcar mas sem expecificar mas parece que não dá
	i = float32(j)
}

////////////////////////////////////////////////////////////

func slicing() {
	//é possível fazer slicing(curtar) arrays

	numSlice := []int{5, 6, 7, 8, 9}

	//numSlice2 guarda os valores que estão no numSlice desde o index 3 até ao index 5-1
	numSlice2 := numSlice[2:5]
	fmt.Println(numSlice, numSlice2)

	//omitir um dos lados irá meter desde o inicio ou no fim dependendo do lado
	fmt.Println(numSlice[:2], numSlice[3:])

	//cria um vetor com um tamanho inicial de 5 com capacidade maxima de 10
	numSlice3 := make([]int, 5, 10)

	copy(numSlice3, numSlice)

}

////////////////////////////////////////////////////////////

// os ... significa que pode receber um indeferido numero de argumentos
func addAll(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue += value
	}
	return finalValue
}

////////////////////////////////////////////////////////////

//para se fazer polomorfismo tem se de fazer da seguinte forma
type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width

}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func testingMetaClasses() {
	r := Rectangle{20, 20}
	c := Circle{4}
	fmt.Println(r, getArea(r))
	fmt.Println(c, getArea(c))
}

////////////////////////////////////////////////////////////

func deferTesting() {
	//defer é uma keyword especial que serve para dizer que a seguinte
	//função tem de correr no final do scope,
	//logo a função print2 irá correr depois de print1
	defer print2()
	print1()
}

func print2() { fmt.Println(2) }
func print1() { fmt.Println(1) }

////////////////////////////////////////////////////////////

//para mudar os valores das variaveis o conceito é o mesmo que em c
func testPointers() {
	i := 2
	fmt.Println("i :", i)
	changeValueOfTo(&i, 4)
	fmt.Println("i :", i)
	ptrY := new(int)
	fmt.Println("ptrY :", *ptrY)
	changeValueOfTo(ptrY, 3)
	fmt.Println("ptrY :", *ptrY)
}

func changeValueOfTo(of *int, to int) {
	*of = to
}

////////////////////////////////////////////////////////////

func playingWithStrings() {
	sampString := "ero, me neme is Joan, the horse"
	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "me"))
	fmt.Println(strings.Count(sampString, "me"))
	fmt.Println(strings.ReplaceAll(sampString, "me", "mo"))
	fmt.Println(strings.Replace(sampString, "e", "i", 3))

	sampString = "1,2,3,4,5,6"
	fmt.Println(strings.Split(sampString, ","))
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters: ", listOfLetters)
	fmt.Println(strings.Join([]string{"3", "2", "1"}, ","))
}

////////////////////////////////////////////////////////////

func fileManipulation() {
	file, err := os.Create("samp.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("This is some random text")
	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")
	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)
	fmt.Println(readString)

}

////////////////////////////////////////////////////////////

func randIntnNoZero(n int) int {
	var value int = 0
	for value == 0 {
		value = rand.Intn(n)
	}
	return value
}

func generateRandomStringOfSize(size int) string {
	if size <= 0 {
		return ""
	}
	//Only lowercase
	charSet := "abcdedfghijklmnopqrst"
	var output strings.Builder
	for i := 0; i < size; i++ {
		output.WriteString(string(charSet[rand.Intn(len(charSet))]))
	}
	return output.String()
}

func randomness() {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Int()
	randFloat := rand.Float64()
	randString1 := generateRandomStringOfSize(randIntnNoZero(10))
	randString2 := generateRandomStringOfSize(randIntnNoZero(10))

	fmt.Println(randInt, randFloat, randString1, randString2)
}

////////////////////////////////////////////////////////////

func httpServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handlerEarth)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Herro Brothers\n")
}

func handlerEarth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Herro Earth Brothers\n")
}

////////////////////////////////////////////////////////////

// em go tens a keyword go que inicia uma função de forma paralela

//wait group para esperar pelas go routines acabarem
var wg sync.WaitGroup

func goRoutines() {

	//no final vai esperar
	defer wg.Wait()
	for i := 0; i < 10; i++ {
		//diz ao grupo para esperar por mais um
		wg.Add(1)
		go count(i, &wg)
	}

	fmt.Println("Todos os processos lançados, a esperar...")
}

func count(id int, wg *sync.WaitGroup) {
	//in the end of the function it will execute what is in defer
	//that is in case of errors or things like that, so that it
	//executes in the end of the scope it is defined that it will be done
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Println(id, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

////////////////////////////////////////////////////////////

//go tem o conceito de channels que são canais de typos básicos
//para transferir informações entre go concurrencies(paralelizar)
//TODO: verificar os problemas de DATA RACING
var pizzaName string = ""
var pizzaNum int = 0

func channels() {
	stringChan := make(chan string)
	defer wg.Wait()
	for i := 0; i < 3; i++ {
		wg.Add(3)
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
	}
}

func makeDough(stringChan chan string) {
	defer wg.Done()
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)
	stringChan <- pizzaName
	fmt.Println("Make Dough and Send for Sauce in", pizzaNum)
}

func addSauce(stringChan chan string) {
	defer wg.Done()
	pizza := <-stringChan
	fmt.Println("Add Sauce and Send", pizza, "for toppings")
}

func addToppings(stringChan chan string) {
	defer wg.Done()
	pizza := <-stringChan
	fmt.Println("Add Toppings to ", pizza, "and ship")

}

////////////////////////////////////////////////////////////

func main() {
	channels()
}
