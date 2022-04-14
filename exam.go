package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Task struct {
	Description string
	Done        bool
}

var task []Task

//On fait une boucle qui va chercher la list des tache a faire en incrémentant i, dans le premier cas ID : 1, task: cuisine
func list(rw http.ResponseWriter, _ *http.Request) {
	var concac string		
	for i := 1; i < len(task); i++ {
		if task.Done == false //si la tache n'est pas done alors on l'affiche sinon on passe a l'autre
		concac = concac + "ID: " + strconv.itoa(i) + ", task: " + task.Description(i) //tentative de boucle pour chaque element de task
	}
	return concac
}
rw.WriteHeader(http.StatusOK)
Write([]byte(concac)) (int, error) //on fait le message pour e client ne transformant la chaine en bytes
}

func slashdone(sd http.ResponseWriter, _ *http.Request) {
	switch req.Method {	
	case http.MethodGet:
		for i := 1; i < len(task); i++ {
			if task.Done == true
			//On doit garder la tache (pas terminé)
	}
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body) 
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		http.Error(rw,"can't read body", http.StatusBadRequest,)
		convert := string(body)
 			return
	default:
		rw.WriteHeader(http.StatusBadRequest)
	}
}
//Fin de mon avnture, compliqué de tester, beaucoup d'erreur de syntaxe mais j'ai avancé un maximum

func slashadd(sa http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {     // On check le type de request, si différent de pos on met un message d'erreur
		rw.WriteHeader(http.StatusBadRequest)
	}
	body, err := ioutil.ReadAll(r.Body) 
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		http.Error(rw,"can't read body", http.StatusBadRequest,)
 return
	}
	convert := string(body) //transformation en string 
	newtask := Task{convert, false} //nouvelle tache avec le body en string + le bool
	task = append(task[], newtask) // On ajoute la tâche 
	rw.WriteHeader(http.StatusOK)
}
func main() {
	fmt.Println()
	http.ListenAndServe("localhost:4000", nil)
	http.HandleFunc("/", list) 
	http.HandleFunc("/done", slashdone)
	http.HandleFunc("/add", slashadd)
}
