package hotel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hotel struct {
	HotelId  int
	Name     string
	BookSlot []bool
}

type Booking struct {
	BookingId int
	HotelId   int
	CheckIn   int
	CheckOut  int
}

func NewHotel(name string, id int) *Hotel {
	newHotel := new(Hotel)

	newHotel.HotelId = id
	newHotel.Name = name
	newHotel.BookSlot = make([]bool, 365)

	return newHotel
}

func NewBooking(bookingId int, hotelId int, checkin int, checkout int) *Booking {
	booking := new(Booking)
	booking.BookingId = bookingId
	booking.HotelId = hotelId
	booking.CheckIn = checkin
	booking.CheckOut = checkout

	return booking
}

func myAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return i
}

func HotelBois() {
	n := 0
	fmt.Scanf("%d", &n)

	hotelIncrementalId := 1
	bookingIncrementalId := 1

	hotels := make(map[int]Hotel)
	bookings := make(map[int]Booking)

	commandIter := 0
mainLoop:
	for commandIter < n {
		in := bufio.NewReader(os.Stdin)

		line, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		line = strings.Trim(line, "\n")
		cmd := strings.Split(line, " ")

		switch cmd[0] {
		case "create":
			hotelName := cmd[1]
			newHotel := NewHotel(hotelName, hotelIncrementalId)
			hotels[hotelIncrementalId] = *newHotel

			hotelIncrementalId++
			fmt.Println("== Hotel Created ==")
			commandIter++
		case "book":
			hotelId := myAtoi(cmd[1])
			checkin := myAtoi(cmd[2])
			checkout := myAtoi(cmd[3])

			for _, booked := range hotels[hotelId].BookSlot[checkin:checkout] {
				// fmt.Println(booked)
				if booked {
					fmt.Println("The Date is already booked")
					commandIter++
					continue mainLoop
				}
			}

			for i := checkin; i < checkout; i++ {
				hotels[hotelId].BookSlot[i] = true
			}

			newBooking := NewBooking(bookingIncrementalId, hotelId, checkin, checkout)
			bookings[bookingIncrementalId] = *newBooking

			bookingIncrementalId++

			fmt.Println("== Booked ==")
			commandIter++
		case "cancel":
			hotelId := myAtoi(cmd[1])
			bookingToDelete := myAtoi(cmd[2])

			theBooking := bookings[bookingToDelete]

			for i := theBooking.CheckIn; i < theBooking.CheckOut; i++ {
				hotels[hotelId].BookSlot[i] = false
			}

			delete(bookings, bookingToDelete)
			commandIter++
		case "exit":
			fmt.Println("== Exit ==")
			break mainLoop
		}
	}

	fmt.Println("== Result ==")

	for _, h := range hotels {
		fmt.Println(h.Name)

		for _, x := range bookings {
			if h.HotelId == x.HotelId {
				fmt.Println("Booking Id", x.BookingId, ":", x.CheckIn, "->", x.CheckOut)
			}
		}

		fmt.Println()
	}
}
