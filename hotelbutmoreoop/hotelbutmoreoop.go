package hotelbutmoreoop

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HotelEntity struct {
	HotelId  int
	Name     string
	BookSlot []bool
}

type BookingEntity struct {
	BookingId int
	HotelId   int
	CheckIn   int
	CheckOut  int
}

type BookingManager struct {
	HotelIncrementalId   int
	BookingIncrementalId int

	Hotels   map[int]*HotelEntity
	Bookings map[int]*BookingEntity
}

func (m *BookingManager) Initialize() {
	m.HotelIncrementalId = 1
	m.BookingIncrementalId = 1

	m.Hotels = make(map[int]*HotelEntity)
	m.Bookings = make(map[int]*BookingEntity)
}

func (m *BookingManager) NewHotel(name string) *HotelEntity {
	newHotel := new(HotelEntity)

	newHotel.HotelId = m.HotelIncrementalId
	newHotel.Name = name
	newHotel.BookSlot = make([]bool, 365)

	m.Hotels[m.HotelIncrementalId] = newHotel
	m.HotelIncrementalId++

	return newHotel
}

func (m *BookingManager) NewBooking(hotelId int, checkin int, checkout int) (*BookingEntity, bool) {
	for _, booked := range m.Hotels[hotelId].BookSlot[checkin:checkout] {
		if booked {
			fmt.Println("The Date is already booked")
			return nil, false
		}
	}

	for i := checkin; i < checkout; i++ {
		m.Hotels[hotelId].BookSlot[i] = true
	}

	booking := new(BookingEntity)
	booking.BookingId = m.BookingIncrementalId
	booking.HotelId = hotelId
	booking.CheckIn = checkin
	booking.CheckOut = checkout

	m.Bookings[m.BookingIncrementalId] = booking
	m.BookingIncrementalId++

	return booking, true
}

func (m *BookingManager) CancelBooking(hotelId int, bookingId int) {
	booking := m.Bookings[bookingId]

	for i := booking.CheckIn; i < booking.CheckOut; i++ {
		m.Hotels[hotelId].BookSlot[i] = false
	}

	delete(m.Bookings, bookingId)
}

func (m *BookingManager) Showdown() {
	for _, h := range m.Hotels {
		fmt.Println("Hotel : ", h.Name)

		for _, b := range m.Bookings {
			if h.HotelId == b.HotelId {
				fmt.Println("Booking Id", b.BookingId, ":", b.CheckIn, "->", b.CheckOut)
			}
		}

		fmt.Println()
	}
}

func myAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return i
}

func HotelBoisWithOOP() {
	n := 0
	fmt.Scanf("%d", &n)

	manager := new(BookingManager)
	manager.Initialize()

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
			manager.NewHotel(hotelName)

			fmt.Println("== Hotel Created ==")
			commandIter++
		case "book":
			hotelId := myAtoi(cmd[1])
			checkin := myAtoi(cmd[2])
			checkout := myAtoi(cmd[3])

			manager.NewBooking(hotelId, checkin, checkout)

			fmt.Println("== Booked ==")
			commandIter++
		case "cancel":
			hotelId := myAtoi(cmd[1])
			bookingToDelete := myAtoi(cmd[2])

			manager.CancelBooking(hotelId, bookingToDelete)
			fmt.Println("== Canceled ==")
			commandIter++
		case "exit":
			fmt.Println("== Exit ==")
			break mainLoop

		}
	}

	fmt.Println("== Result ==")
	manager.Showdown()
}
