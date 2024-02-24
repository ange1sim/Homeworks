package operations

import (
	"hw/models"
	"fmt"
	"log"
	"math/rand"
)

func TakeOrder(client *models.Patient) *models.Patient {

	//Проверка клиента на возраст
	if client.Age <= 15 {
		fmt.Println("Слишком мал для операции!")
		return &models.Patient{}
	} else if client.Age >= 150 {
		fmt.Println("Слишком...стар?")
	}

	//Клиент получает случайный билетик с номером.
	client.Order = rand.Intn(20) + 1
	return client
}

func BeginOperation(client *models.Patient) *models.Patient {
	var (
		//Количество людей в очереди + очередь клиента. Сделано для того, чтобы очередь никогда не становилась меньше, чем очередь клиента.
		peoples int = rand.Intn(20) + client.Order
	)
	//Раньше тут был бесконечный цикл с рандомайзером, но я решил показать функционал циклов на деле без infinite loop более практично, чем на прошлом дне. Так гораздо лучше, так как очередь теперь реальная.
	for i := 1; i <= peoples; i++ {

		//УСТАРЕЛО //Симулируем ожидание в очереди с помощью rand (до 20)
		// i := rand.Intn(21) + 1

		switch i {
		//Если число совпало с очередным номером клиента, ему пишется диагноз, даётся выписка и его отправляют на операцию (функция Operation)
		case client.Order:
			log.Println("Дождался!", i)
			client.MedBook.Annotation = true
			client.MedBook.Diagnose = "Установка нового сердца"
			client = Operation(client)
			return client
		//Очередь клиента не подошла, но он считает, сколько человек прошло.
		default:
			log.Println("Очередной человек прошёл мимо: ", i)
		}
	}
	return client
}

func Operation(client *models.Patient) *models.Patient {
	switch {
	//У клиента нет выписки
	case !client.MedBook.Annotation:
		fmt.Println("Уходи, без выписки пробрался!")
		return client
	//У клиента не тот диагноз
	case client.MedBook.Diagnose != "Установка нового сердца":
		fmt.Println("Не та операция")
		return client
	default:
		//У Операция началась, симулируем ненастоящую операцию с дурным рандомом по удаче, как в настольной игре с кубиком
		luck := rand.Intn(4)
		switch luck {
		case 0:
			fmt.Println("Ой... а где нога?")
			client.Organs.Leg--
		case 1:
			fmt.Println("Не повезло! Рука пропала!")
			client.Organs.Hands--
		case 2:
			fmt.Println("А у него всегда была одна почка?")
			client.Organs.Livers--
		case 3:
			//Успешная операция
			fmt.Println("Первый человек с новым сердцем!")
			client.Organs.Heart = true
			client.MedBook.Annotation = false
			client.MedBook.Diagnose = "Живой!"
		}
	}
	return client
}
