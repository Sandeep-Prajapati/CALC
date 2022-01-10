package calctest

import (
	"context"
	"fmt"
	"testing"

	pb "example.com/sandeepkp/calc/proto"
	"google.golang.org/grpc"
)

var client pb.CalcServiceClient

func getConnection() pb.CalcServiceClient {
	if client == nil {
		address := "127.0.0.1:8080"

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			panic(fmt.Sprintf("Unable to connect %s", err))
		}
		// defer conn.Close()
		fmt.Println("Starting client on", address)
		client = pb.NewCalcServiceClient(conn)
	}
	return client
}

func TestAdd(t *testing.T) {
	c := getConnection()
	fmt.Printf("Address %p\n", c)

	var addTests = []struct {
		a        int64 // input
		b        int64 // input
		expected int64 // expected result
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
		{2, 5, 7},
		{-2, -3, -5},
		{-1, 8, 7},
		{4, -2, 2},
	}

	// First call
	for _, params := range addTests {
		request := pb.Request{A: params.a, B: params.b}
		response, err := c.Add(context.Background(), &request)
		if err != nil {
			t.Error(err)
			continue
		}
		if response.Result != params.expected {
			t.Errorf("c.Add(%d,%d): expected: %d, actual: %d", params.a, params.b, params.expected, response.Result)
		}
	}

}

func TestSub(t *testing.T) {
	c := getConnection()
	fmt.Printf("Address %p\n", c)

	var addTests = []struct {
		a        int64 // input
		b        int64 // input
		expected int64 // expected result
	}{
		{0, 0, 0},
		{0, 1, -1},
		{1, 0, 1},
		{2, 5, -3},
		{-2, -3, 1},
		{-1, 8, -9},
		{4, -2, 6},
	}

	// Second call
	for _, params := range addTests {
		request := pb.Request{A: params.a, B: params.b}
		response, err := c.Sub(context.Background(), &request)
		if err != nil {
			t.Error(err)
			continue
		}
		if response.Result != params.expected {
			t.Errorf("c.Sub(%d,%d): expected: %d, actual: %d", params.a, params.b, params.expected, response.Result)
		}
	}

}

func TestMul(t *testing.T) {
	c := getConnection()
	fmt.Printf("Address %p\n", c)

	var addTests = []struct {
		a        int64 // input
		b        int64 // input
		expected int64 // expected result
	}{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{2, 5, 10},
		{-2, -3, 6},
		{-1, 8, -8},
		{4, -2, -8},
	}

	// Third call
	for _, params := range addTests {
		request := pb.Request{A: params.a, B: params.b}
		response, err := c.Mul(context.Background(), &request)
		if err != nil {
			t.Error(err)
			continue
		}
		if response.Result != params.expected {
			t.Errorf("c.Mul(%d,%d): expected: %d, actual: %d", params.a, params.b, params.expected, response.Result)
		}
	}

}

func TestDiv(t *testing.T) {
	c := getConnection()
	fmt.Printf("Address %p\n", c)

	var addTests = []struct {
		a        int64 // input
		b        int64 // input
		expected int64 // expected result
	}{
		// {0, 0, 0},
		{0, 1, 0},
		{1, 0, 1},
		{2, 5, 0},
		{-2, -3, 0},
		{-1, 8, 0},
		{4, -2, -2},
	}

	// Fourth call
	for _, params := range addTests {
		request := pb.Request{A: params.a, B: params.b}
		response, err := c.Div(context.Background(), &request)
		if err != nil {
			t.Errorf("%s, params:%d, %d", err, params.a, params.b)
			continue
		}
		if response.Result != params.expected {
			t.Errorf("c.Div(%d,%d): expected: %d, actual: %d", params.a, params.b, params.expected, response.Result)
		}
	}
}
