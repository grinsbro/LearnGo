package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	// runtime - это встроенная библиотека в Go, которая позволяет управлять потоками
	// fmt.Println(runtime.NumCPU()) // Здесь выводится количество одновременных выполнений горутин, которое поддерживает компьютер. То есть невозможно выполнять горутин больше этого числа на данном компьютере чисто физически
	// runtime.GOMAXPROCS(2)         // Так задается вручную, сколько одновременно горутин можно выполнять. Если поставить число больше количества ядер на компьютере, то все равно будет выполняться максимально возможное количество

	// go printInt(10) // Таким образом создается горутина. Горутина - это легковесный поток (горутина ~2Кб, а обычный поток ~1Мб)
	// Горутины могут работать как конкурентно, так и параллельно
	// Отличие горутин от потоков также в том, что ими управляет планировщик Go, а не планировщик ОС
	// Треды переключается автоматически под капотом Go
	// При такой записи будет выводиться сразу Exit..., потому что как только завершается главная горутина main, мы выходим из программы и созданная вручную горутина даже не успевает отработать

	// runtime.Gosched() // Этот метод позволяет вручную переключиться между горутинами
	// Теперь функция выполнится полностью, а после завершения вернется в главную горутину и завершит ее выполнение
	// time.Sleep(time.Second) Также например можно так сказать планирощику Go перейти к другому потоку, потому что он поймет, что можно выполнять дргуую горутину пока в этой действует sleep
	// fmt.Println("Exit...")

	// withWait()

	// writeWithMutex()

	// nilChannel()

	// unbufferedChannel()

	// bufferedChan()

	// baseSelect()

	// baseKnowledge()

	// workerPool()

	// chanPromise()

	// withErrGroup()

	// addAtomic()

	// storeLoadSwap()

	compareAndSwap()
}

// func printInt(n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Println(i)
// 	}
// }

// func withWait() {
// 	var wg sync.WaitGroup // Так создается wait группа, она нужна для того, чтобы синхронизировать горутины

// 	for i := 0; i <= 10; i++ {
// 		wg.Add(1) // Метод Add позволяет добавить таску в счетчик wait группы

// 		go func(i int) {
// 			fmt.Println(i)
// 			wg.Done() // Этот метод говорит wait групе, что это таска выполнена и можно убирать ее из счетчика
// 		}(i)
// 	}

// 	wg.Wait() // Этот метод говорит, что основная горутина не должна быть завершена пока в wait группе еще есть горутины для выполнения
// 	fmt.Println("Exit...")
// }

// func writeWithMutex() {
// 	start := time.Now()
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex // Так создается Mutex. Он дает возможность предоставить эксклюзивные права к какому-либо участку кода только одной горутине
// 	var counter int

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			time.Sleep(time.Nanosecond)

// 			mu.Lock() // Так локается участок кода. Остальные горутины пока не могут получить доступ к этому участку висят в ожидании
// 			counter++
// 			mu.Unlock() // Так анлокается участок кода и доступ к нему получает опять только одна горутина
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(counter)
// 	fmt.Println(time.Now().Sub(start).Seconds())
// }

// Каналы в Go
// Каналы нужны для безопасной передачи данных между горутинами
// Каналы являются более высокоуровневой структурой, потому что они отбрасывают необходимость лочить и разлочить что-то вручную
// Под капотом каналы выглядят так:
// type chan struct {
// 	mx sync.Mutex
// 	buffer []T
// 	readers []Goroutines
// 	writes []Goroutines
// }

// Применение каналов
func nilChannel() {
	var nilChannel chan int // Каналы объявляются с помощью ключевого слова chan, а затем пишется тип данных, который будет передаваться между горутинами
	// Это nil канал, то есть длина его буфера и capacity равны нулю
	// Если попытаться записать данные в такой канал или прочитать, то произойдет deadlock, потому что у nil канала нет буфера
	fmt.Printf("Len: %d Cap: %d\n", len(nilChannel), cap(nilChannel))

	// Таким образом записываеются данные в канал:
	// nilChannel <- 1 // Запись происходит с использованием оператора <-. Сейчас этот код вызовет deadlock

	// Таким образом можно прочитать данные из канала
	// <-nilChannel // Также используется оператор <-, но уже слева от самого канала. Этот код в данном случае тоже приведет к deadlock

	// Также каналы можно закрывать:
	// close(nilChannel) // Но в данном случае также выпадет паника, потому что это nil канал
}

func unbufferedChannel() {
	unbufferedChannel := make(chan int) // Так объявляется небуфиризированный канал
	// Длина и capacity такого канала также будут равны нулю, но с ним уже можно работать

	// При работе с каналом такого типа обязательно нужно, чтобы было две горутины, которые занимают очередь на чтение и запись, потому что иначе выпадет паника и deadlock
	// unbufferedChannel <- 1
	// <- unbufferedChannel

	go func(chanForWriting chan<- int) {
		time.Sleep(time.Second)
		unbufferedChannel <- 1
	}(unbufferedChannel)

	value := <-unbufferedChannel
	fmt.Println(value)

	go func(cnahForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-unbufferedChannel
		fmt.Println(value)
	}(unbufferedChannel)

	unbufferedChannel <- 2

	// Такой канал можно закрыть и не выпадет паника.
	// Но если попытаться записать что-то в закрытый канал, то выпадет паника
	// Также если попытаться закрыть закрытый канал, то тоже выпадет паника.

	// У каналов есть направленность(Только на чтение или запись)
	// Чтобы строго задать направленность канала, можно при его объявлении написать:
	// unBufferredChan := make(chan<- int) - Канал только для записи
	// unBufferredChan := make(<-chan int) - Канал только для чтения
	// Также можно задать строгую направленность в горутинах, которые используют канал, чтобы сразу была понятна направленность горутины
}

func bufferedChan() {
	bufferedChan := make(chan int, 2) // Таким образом объявляется буферизированный канал. У него есть capacity, который был передан вторым аргументом

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChan), cap(bufferedChan))

	// Главное отличие буферизированного канала от небуферизированного в том, что если попытаться записать данные в канал без явной горутины, которая читает эти данные, то паника не выпадет
	// Но только пока буфер не заполнен. Если попытаться записать в канал при полном буфере еще что-то и не указать горутину на чтение, то будет deadlock
	// Также это работает и с чтением, но только deadlock будет, если буфер пустой и нет в очереди горутины на запись
	bufferedChan <- 1
	bufferedChan <- 2

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChan), cap(bufferedChan))

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	fmt.Printf("Len: %d Cap: %d\n", len(bufferedChan), cap(bufferedChan))

}

// Select в работе с каналами

func baseSelect() {
	buffChannel := make(chan string, 1)

	buffChannel <- "first"

	select { // Select в работе с каналами это то же, что и switch case.
	// Select распознает три типа операций, блокирующие, неблокирующие и дефолтные
	case str := <-buffChannel: // В данном случае это неблокирующая операция и будет выполнена она
		fmt.Println("read", str)
	case buffChannel <- "second": // Это в данном случае блокирующая операци
		fmt.Println("write", <-buffChannel, <-buffChannel)
	}
	// Под капотом Select смотрит какой из кейсов неблокирующий и вызывает его
	// Если, например, оба кейса неблокирующие, то select вызовет один из них рандомно
	// Если же кейсы только блокирующие и их невозможно выполнить сразу же, то вызывается ветка default

	timer := time.After(time.Second) // time.After это канал, данные в котором появляются по истечении срока, переданного в аргументе
	// Здесь я выношу его в отдельную переменную, потому что хочу использовать таймер в цикле for, и если не сделать так, то таймер на каждой итерации будет обновляться и никогда не истечет

	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		for i := 0; i <= 1000; i++ {
			select {
			case <-timer:
				fmt.Println("Время вышло")
			default:
				resultChan <- i
			}
		}
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}

// Контекст

func baseKnowledge() {
	// Так создается контекст. Контекст нужен, чтобы хранить какие-тоо базовые значения
	// Например, куки, данные о пользователе и тд
	ctx := context.Background() // Это бэкграунд контекст

	// toDo := context.TODO() // Такой тип контекста нужен преимущественно для тестирования или для проектирования в функцниях, где будет использоваться контекст

	ctxWithValue := context.WithValue(ctx, "name", "Ilia") // Таким образом можно положить данные в контекст. Нужен родительский контекст, ключ и значение
	fmt.Println(ctxWithValue.Value("name"))                // Чтобы получить данные из контекста используется функция Value, куда передается ключ

	// !!!
	// Считается антипаттерном складывать все данные в контекст и затем вызывать их таким образом
	// !!!

	withCancel, cancel := context.WithCancel(ctx) // Контекст также может сообщать о заврешении какой-либо задачи. Для этого нужно, чтобы контекст мог быть отменен
	// Чтобы записать контекст с возможностью отмены нужно вызвать метод WithCancel(), который на основе родительского возвращает новый контекст и функцию для отмены
	fmt.Println(withCancel.Err()) // Сейчас вывод будет nil
	cancel()                      // Вызываю функцию отмены, которая записалась в переменную cancel
	fmt.Println(withCancel.Err()) // Теперь вывод будет context canceled

	// !!!
	/*
		Считается плохой практикой отменять контекст в каких-либо функциях или на других уровнях отличных от того, где была создана функция отмены
		Лучше всегда закрывать контекст на одном уровне с его созданием
	*/
	// !!!

	// Контекст с возможностью отмены также можно создать с каким-то дедлайном или таймаутом
	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*5)) // Дедлайн устанавливается на определенное время. То есть нужно получить время сейчас и добавить к нему желаемое время таймаута
	defer cancel()
	fmt.Println(withDeadline.Err())
	fmt.Println(<-withDeadline.Done())

	// Но лучше работать с контекстом с таймаутом
	withTimeout, cancel := context.WithTimeout(ctx, time.Second*5) // В это случае нужно просто передать через сколько наступить таймаут без необходимости передачи времени на момент создания контекста
	defer cancel()
	fmt.Println(<-withTimeout.Done())

}

// Паттерн использования контекста:

func workerPool() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	numbersToProcess, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i <= runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess, processedNumbers)
		}()
	}

	go func() {
		for i := 0; i <= 1000; i++ {
			numbersToProcess <- i
		}
		close(numbersToProcess)
	}()

	go func() {
		wg.Wait()
		close(processedNumbers)
	}()

	var counter int

	for resultValue := range processedNumbers {
		counter++
		fmt.Println(resultValue)
	}

	fmt.Println(counter)

}

func worker(ctx context.Context, toProcess <-chan int, processed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			processed <- value * value
		}
	}
}

// Каналы как промисы
// Имитация онлайн запроса:
func makeRequest(num int) <-chan string {
	resultChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		resultChan <- fmt.Sprintf("Запрос номер: %d", num)
	}()
	return resultChan
}

// Теперь делаем "запросы"
func chanPromise() {
	firstResponse := makeRequest(1)
	secondResponse := makeRequest(2)

	// Выполняем какой-то другой код...
	fmt.Println("Что-то происходит...")

	fmt.Println(<-firstResponse, <-secondResponse)
}

// Таким образом получается, что каналы могут работать как промисы в JavaScript, то есть мы можем сделать какой-либо запрос, но код все равно продолжит исполняться. А когда нам понадобится резльтат запроса, мы его выведем

// Error Group

func withErrGroup() {
	errG, ctx := errgroup.WithContext(context.Background()) // Error Group позволяет отлавливать ошибки в горутинах и удалять контекст, если возникла какая-либо ошибка
	// У Error group есть два метода Go и Wait. Первый позволяет выполнять горутины и отлавливать ошибки. Метод Go принимает в себя анонимную функцию, которая возвращает ошибку. Метод Wait возвращает ошибку и не принимает аргументов
	errG.Go(func() error {
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
		default:
			fmt.Println("Первая задача")
			time.Sleep(time.Second)
		}
		return nil // Можно записывать возврат не внутри select а после выхода из него. Результат будет одинаковым
	})

	errG.Go(func() error {
		fmt.Println("Вторая задача")
		return fmt.Errorf("непредвиденная ошибка во второй задаче") // Имитирую ошибку во время второй горутины
	})

	errG.Go(func() error {
		select {
		case <-ctx.Done():
		default:
			fmt.Println("Третья задача")
		}
		return nil
	})

	if err := errG.Wait(); err != nil {
		fmt.Println(err)
	}
}

/*
Error Group позволяет отменять контекст, если в результате выполнения одной из горутин выпала ошибка
То есть если создавать ошибку как переменную вручную и в случае ее возникновения записывать данные в каждой из горутин, то они все будут записывать ошибку в ожно и то же место и не будут останавливаться
Error Group же позволяет отменить контекст и не выполнять другие горутины, если одна вернула ошибку
Вообще Error Group это больше синтаксический сахар, чем что-то жизненно необходимое
*/

// Пакет atomic
/*
Данный пакет представляет собой более низкоуровневую реализацию методов, которые записаны в пакете sync, частью которого он является.
Atomic производит вычисления на самом низком уровне, что делает его более высокопроизводительным по сравнению с mutex.
Но стоит отметить, что atomic позволяет выполнять только одно определенгое действие:
AddT
LoadT
StoreT
SwapT
CompareAndSwapT
В то время как в mutex можно записать более сложную логику
Примеры:
*/

func addAtomic() {
	start := time.Now()

	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1) // Метод AddT принимает два аргумента - ссылку на переменную, значение которой нужно изменить и дельту, на которую нужно поменять значение
		}()
	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println("Время с атомик: ", time.Since(start).Seconds())
}

func storeLoadSwap() {
	var counter int64

	fmt.Println(atomic.LoadInt64(&counter)) // Метод LoadT позволяет получить значение из переменной

	atomic.StoreInt64(&counter, 5) // Метод StoreT позволяет поместить какое-либо значение в переменную, но только того типа, который подходит
	fmt.Println(atomic.LoadInt64(&counter))

	fmt.Println(atomic.SwapInt64(&counter, 10)) // Метод SwapT позволяет поменять значение в переменной на то, которое нужно, но он возвращает старое значение
	fmt.Println(atomic.LoadInt64(&counter))
}

func compareAndSwap() {

	var (
		counter int64
		wg      sync.WaitGroup
	)

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) { // Метод CompareAndSwapT возвращает булевое значение и меняет одно значение на другое только если находит значение переданное вторым аргументом
				return
			}

			fmt.Println("Процесс завершен горутиной номер:", i)
		}(i)
	}

	wg.Wait()
	fmt.Println(counter)

}
