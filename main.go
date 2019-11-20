package main // All executable code needs a main package
import (
	"log"
	"net/http" // Import log for logging errors, net/http as we need to make api requests.

	"github.com/gorilla/mux" // We install Gorilla Mux so we can handle path and query parameters. 'net/http' does not allow us to do this.
	_ "github.com/lib/pq"
)

type server struct{}

//Location struct (like the model from JS MVC). A struct is like a JS Class
// type Location struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	Text string `json:"text"`
// }
//Create an array-like object (arrays in Go must declare how many elements they contain. For arrays of unknown quantity, we use 'slice', []. We are telling it it will contain some Location structs)
//var locations []Location
// WE previously has a struct (object-type thing in Go) called 'server' below. This was so we could make a HTTP request
func get(w http.ResponseWriter, r *http.Request) { //http.Request is the client request. http.ResponseWriter assembles the servers response. By writing to it, we send data to the client.
	// 'w' is a struct with properties and methods related to the response that is sent to the client. It let's us define certain things in reponse like header, body, status code etc. 'r' struct stores values received from client. This includes things like request method ('GET' etc), headers, bodies etc. the '*' before http.Request indicates it is a pointer to it, as the r struct is very large so copying it is resource expensive.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "GET called"}`))
}
func apiHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "GET to API called"}`))
}
func getLocations(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "This would be a json object of all our locations"}`))
	//json.NewEncoder(w).Encode(locations) // This returns an array of location objects that was initialised in the 'main' function. Encode also takes care of the status code!!
}
func getSingleLocation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "I am a single location"}`))
	// params := mux.Vars(r)            // This statement allows us to take parameters from the URL request
	// for _, item := range locations { //range specifies anything that is iterable. This line is similar to a For Loop
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(&Location{})
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "not found"}`))
}

// The above functions are like the Controllers in our JS backend
func main() {
	// locations = append(locations, Location{ID: "1", Name: "Hacienda", Text: "I am the Hacienda"})
	// locations = append(locations, Location{ID: "2", Name: "The Twisted Wheel", Text: "I am the Twisted Wheel"})
	r := mux.NewRouter() // We use this to enable Gorilla Mux. With 'r', we are creating a router object.
	r.HandleFunc("/", get).Methods("GET")
	r.HandleFunc("/api", apiHandle).Methods("GET")
	r.HandleFunc("/api/locations", getLocations).Methods("GET")
	r.HandleFunc("/api/locations/{id}", getSingleLocation).Methods("GET")
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":9090", r)) // This is Go's vrsion of app.listen
}

// The above func is liek the Router we use in JS backend
