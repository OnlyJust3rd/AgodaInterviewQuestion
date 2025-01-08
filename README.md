# Hotel Management System - Command Line Program

## Problem Statement

**Time:** 1 hour

Create a command line program that works as a simple hotel management system. It should process `n` transactions. Each transaction can either:

- **Create** hotels
- **Book** a hotel within a specified date range
- **Cancel** bookings

The program should handle potential conflicts in booking dates that might occur from user input. When the program terminates, it should **report all bookings** for all hotels.

### Requirements

- Hotels and bookings should have auto-incremented IDs to identify them. Users don't need to input IDs manually.
- Booking IDs should not rely on specific hotels. They should be **system-wide counting numbers**.
- **Assume:**
  - There is **only one room per hotel**.
  - Dates are represented as numbers from **1 to 365**.

---

## Input Format

- The first line contains an integer `n` (where `1 ≤ n ≤ 100,000,000`), representing the number of commands.
- The next `n` lines describe each transaction. The transaction can follow one of these formats:

1. `create hotel <hotel-name>`
2. `book <hotel-id> <check-in-date> <check-out-date>`
3. `cancel <hotel-id> <booking-id>`

---

## Output Format

- Report all hotels in the system with their bookings. (Any format or order is acceptable.)

---

## Sample 1

### Input

```
7
create hotel A
book 1 5 10
create hotel B
book 2 1 10
book 1 12 18
book 2 20 25
cancel 2 4
```

### Output

```
Hotel: A
Booking Id 1: 5 -> 10
Booking Id 3: 12 -> 18

Hotel: B
Booking Id 2: 1 -> 10
```

**Explanation:**

- `Hotel A` is created first and automatically gets hotel ID `1`.
- A booking is made at hotel A from day `5` to `10` (assigned booking ID `1`).
- `Hotel B` is created next (hotel ID `2`).
- A booking is made at hotel B from day `1` to `10` (assigned booking ID `2`).
- Another booking is made at hotel A from day `12` to `18` (assigned booking ID `3`).
- A booking is made at hotel B from day `20` to `25` (assigned booking ID `4`).
- The command `cancel 2 4` removes booking ID `4` from hotel B.

---

## Sample 2

### Input

```
5
create hotel Test
book 1 1 5
book 1 10 14
book 1 2 4
book 1 16 20
```

### Output

```
Hotel: Test
Booking Id 1: 1 -> 5
Booking Id 2: 10 -> 14
Booking Id 3: 16 -> 20
```

**Explanation:**

- `Hotel Test` is created and gets hotel ID `1`.
- Booking from day `1` to `5` is assigned booking ID `1`.
- Booking from day `10` to `14` is assigned booking ID `2`.
- The booking from day `2` to `4` conflicts with booking ID `1` and is not created.
- Booking from day `16` to `20` is successfully created and assigned booking ID `3`.

---

## Marking Criteria

- This problem evaluates **Object-Oriented programming skills** and **code readability**.
- We want to see if you can understand the problem and create objects that reflect the required behavior.

---

## Tips

- Focus on **class design** and **interface functionality** first. Work on implementation, input parsing, and result display later.
- Use **meaningful names** for classes, variables, and methods.
- **Handle edge cases** to improve your solution.
- If in doubt, ask for clarification.
