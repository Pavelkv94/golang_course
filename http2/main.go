
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

type PaymentInfo struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Time        time.Time `json:"time"`
}

func (p PaymentInfo) Validate() bool {
	if p.USD < 0 {
		return false
	}
	if p.FullName == "" {
		return false
	}
	if p.Address == "" {
		return false
	}

	return true
}

type HttpResponse struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

var money = 1000
var paymentHistory = make([]PaymentInfo, 0)
var mtx = sync.Mutex{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	 filter := r.URL.Query().Get("filter") //! query params
	 fmt.Println("filter: ", filter)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	var paymentInfo PaymentInfo
	if err = json.Unmarshal(httpRequestBody, &paymentInfo); err != nil {
		fmt.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: " + err.Error() + " - " + string(httpRequestBody)))
		return
	}

	if !paymentInfo.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: invalid payment info"))
		return
	}

	paymentInfo.Time = time.Now()

	mtx.Lock()
	defer mtx.Unlock()

	if money-paymentInfo.USD < 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: not enough money"))
		return
	}

	money -= paymentInfo.USD
	paymentHistory = append(paymentHistory, paymentInfo)
	pp.Println("paymentHistory: ", paymentHistory)
	fmt.Println("money: ", money)

	httpResponse := HttpResponse{
		StatusCode: http.StatusOK,
		Message: "Payment successful",
	}

	jsonResponse, err := json.MarshalIndent(httpResponse, "", "  ") //MarshalIndent для красивого вывода json(с отступами)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error: " + err.Error()))
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
}

func main_http2() {
	http.HandleFunc("/pay", payHandler)

	// переменная err исчезает после выполнения блока
	if err := http.ListenAndServe(":5000", nil); err != nil {
		fmt.Println("Error: ", err.Error())
	}

	// то же самое что и выше но по другому,
	// в этом случае переменная err остается в области видимости после выполнения блока
	// if err != nil {
	// 	fmt.Println("Error: ", err.Error())
	// }

}
