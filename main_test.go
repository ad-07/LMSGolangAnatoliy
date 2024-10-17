// // 5
package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		// имя теста
		name string
		// значения на вход тестируемой функции
		str []byte
		// желаемый результат
		want   int
		wantEr error
	}{
		// тестовые данные № 1
		{
			name:   "upper vowels",
			str:    []byte("Hello, Go!"),
			want:   10,
			wantEr: nil,
		},
		// тестовые данные № 2
		{
			name:   "mixed vowels",
			str:    []byte("Byte array example"),
			want:   18,
			wantEr: nil,
		},
		{
			name:   "no vowels",
			str:    []byte{0xff, 0xfe, 0xfd},
			want:   0,
			wantEr: ErrInvalidUTF8,
		},
	}
	// перебор всех тестов
	for _, tc := range cases {
		tc := tc
		// запуск отдельного теста
		t.Run(tc.name, func(t *testing.T) {
			// тестируем функцию Sum
			got, gotEr := GetUTFLength(tc.str)
			// проверим полученное значение
			if got != tc.want || gotEr != tc.wantEr {
				t.Errorf("GetUTFLength(%v) = %v, %v; want %v, %v", tc.str, got, gotEr, tc.want, tc.wantEr)
			}
		})
	}
}
