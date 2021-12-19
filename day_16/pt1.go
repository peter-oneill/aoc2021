// Run as `go run pt2.go "$(cat input.txt )"``
package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type packet struct {
	p_version     int
	p_type        int
	p_lit_val     int
	p_sub_packets []packet
}

func main() {
	input := os.Args[1]
	binary_string := input_to_bin_str(input)
	fmt.Println("got input binary: ", binary_string)

	var packets []packet
	var next_pkt_ix = 0

	for {
		new_packet, read_pkt_len := read_packet(binary_string)

		if read_pkt_len == 0 || read_pkt_len == len(binary_string) {
			break
		}
		binary_string = binary_string[read_pkt_len:]
		next_pkt_ix += read_pkt_len
		packets = append(packets, new_packet)
	}

	fmt.Println("packets: ", packets)

	version_sum := sum_pkt_vers(packets)

	fmt.Println("Version sum: ", version_sum)

	return
}

func sum_pkt_vers(packets []packet) int {
	var sum = 0
	for _, pkt := range packets {
		sum += pkt.p_version
		sum += sum_pkt_vers(pkt.p_sub_packets)
	}
	return sum
}

func read_packet(binary_string string) (packet, int) {
	if len(binary_string) < 11 {
		return packet{}, 0
	}

	p_version := bin_str_to_int(binary_string[0:3])
	p_type := bin_str_to_int(binary_string[3:6])
	var p_lit_val = -1
	var next_pkt_ix = 6
	var sub_packets []packet = make([]packet, 0)
	var read_len = 0

	if p_type == 4 {
		p_lit_val, read_len = read_literal_packet(binary_string[6:])

	} else {
		sub_packets, read_len = read_compound_packet(binary_string[6:])
	}
	next_pkt_ix += read_len

	new_packet := packet{p_version, p_type, p_lit_val, sub_packets}
	return new_packet, next_pkt_ix
}

func read_compound_packet(binary_string string) ([]packet, int) {
	var sub_packets []packet = make([]packet, 0)
	var next_pkt_ix = 0

	switch binary_string[0] {
	case '0':
		if len(binary_string) < 16 {
			return nil, 0
		}
		bit_length := bin_str_to_int(binary_string[1:16])
		next_pkt_ix += 16
		if bit_length == 0 {
			return nil, next_pkt_ix
		}
		binary_string = binary_string[next_pkt_ix:]
		for {
			new_packet, read_len := read_packet(binary_string)
			if read_len == 0 {
				fmt.Println("Err unexpected end inside compount packet")
				break
			}
			next_pkt_ix += read_len
			binary_string = binary_string[read_len:]
			sub_packets = append(sub_packets, new_packet)

			if next_pkt_ix >= bit_length {
				break
			}
		}
	case '1':
		num_packets := bin_str_to_int(binary_string[1:12])
		next_pkt_ix += 12
		binary_string = binary_string[next_pkt_ix:]
		var pkt_count = 0
		for {
			new_packet, read_len := read_packet(binary_string)
			next_pkt_ix += read_len
			binary_string = binary_string[read_len:]
			sub_packets = append(sub_packets, new_packet)

			pkt_count++
			if pkt_count >= num_packets {
				break
			}
		}

	}

	return sub_packets, next_pkt_ix
}

func read_literal_packet(binary_string string) (int, int) {
	var buf bytes.Buffer
	var char_ix = 0

	for {
		if binary_string[char_ix] == '0' {
			break
		}

		buf.WriteString(binary_string[char_ix+1 : char_ix+5])
		char_ix += 5
	}

	buf.WriteString(binary_string[char_ix+1 : char_ix+5])
	output_val := bin_str_to_int(buf.String())
	return output_val, char_ix + 5
}

func input_to_bin_str(input string) string {
	var buf bytes.Buffer

	for _, char := range input {
		append_str := hex_char_to_bin_str(char)
		buf.WriteString(append_str)
	}

	return buf.String()
}

func hex_char_to_bin_str(char rune) string {
	switch char {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	default:
		return ""
	}
}

func bin_str_to_int(binary_string string) int {
	var output_val int = 0
	for _, bit := range binary_string {
		output_val *= 2
		int_val, _ := strconv.Atoi(string(bit))
		output_val += int_val
	}
	return output_val
}
