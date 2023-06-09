package trello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBoard(t *testing.T) {
	t.Run("invalid URL", func(t *testing.T) {
		origURL := CreateBoardURL
		CreateBoardURL = "testInvalidURL%s"
		defer func() {
			CreateBoardURL = origURL
		}()

		result, err := CreateBoard(DefaultBoardName)
		assert.Equal(t, "", result)
		assert.NotNil(t, err.Error())
	})

	t.Run("unsupported protocol", func(t *testing.T) {
		origURL := CreateBoardURL
		CreateBoardURL = "testInvalidURL"
		defer func() {
			CreateBoardURL = origURL
		}()

		result, err := CreateBoard(DefaultBoardName)
		assert.Equal(t, "", result)
		assert.NotNil(t, err.Error())
	})

	t.Run("empty response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		origURL := CreateBoardURL
		CreateBoardURL = ts.URL
		defer func() {
			CreateBoardURL = origURL
		}()

		result, err := CreateBoard(DefaultBoardName)
		assert.Equal(t, "", result)
		assert.NotNil(t, err)
	})

	t.Run("error - unauthorized board creation", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer ts.Close()

		origURL := CreateBoardURL
		CreateBoardURL = ts.URL
		defer func() {
			CreateBoardURL = origURL
		}()

		result, err := CreateBoard(DefaultBoardName)
		assert.Equal(t, "", result)
		assert.Equal(t, "failed to create board - status code: 401", err.Error())
	})

	t.Run("successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"id": "abc123a36eaf8d75e160000f"}`))
			assert.Nil(t, err)
		}))
		defer ts.Close()

		origURL := CreateBoardURL
		CreateBoardURL = ts.URL
		defer func() {
			CreateBoardURL = origURL
		}()

		result, err := CreateBoard(DefaultBoardName)
		assert.Equal(t, "abc123a36eaf8d75e160000f", result)
		assert.Nil(t, err)
	})
}

func TestCreateList(t *testing.T) {
	t.Run("invalid URL", func(t *testing.T) {
		origURL := CreateListURL
		CreateListURL = "testInvalidURL"
		defer func() {
			CreateListURL = origURL
		}()

		result, err := CreateList("abc123a36eaf8d75e160000f", "sample list unauthorized", "1")
		assert.Equal(t, "", result)
		assert.NotNil(t, err.Error())
	})

	t.Run("unsupported protocol", func(t *testing.T) {
		origURL := CreateListURL
		CreateListURL = "testInvalidURL%s"
		defer func() {
			CreateListURL = origURL
		}()

		result, err := CreateList("abc123a36eaf8d75e160000f", "sample list unauthorized", "1")
		assert.Equal(t, "", result)
		assert.NotNil(t, err.Error())
	})

	t.Run("empty response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		ts.URL = ts.URL + "/%s"
		origURL := CreateListURL
		CreateListURL = ts.URL
		defer func() {
			CreateListURL = origURL
		}()

		result, err := CreateList("abc123a36eaf8d75e160000f", "sample list unauthorized", "1")
		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("error - unauthorized list creation", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer ts.Close()

		ts.URL = ts.URL + "/%s"
		origURL := CreateListURL
		CreateListURL = ts.URL
		defer func() {
			CreateListURL = origURL
		}()

		result, err := CreateList("abc123a36eaf8d75e160000f", "sample list unauthorized", "1")
		assert.Equal(t, "", result)
		assert.Equal(t, "failed to create list - status code: 401", err.Error())
	})

	t.Run("successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"id": "abc123a36ech8d75e160000f"}`))
			assert.Nil(t, err)
		}))
		defer ts.Close()

		ts.URL = ts.URL + "/%s"
		origURL := CreateListURL
		CreateListURL = ts.URL
		defer func() {
			CreateListURL = origURL
		}()

		result, err := CreateList("abc123a36eaf8d75e160000f", "sample list", "1")
		assert.Equal(t, "abc123a36ech8d75e160000f", result)
		assert.Nil(t, err)
	})
}

func TestCreateCard(t *testing.T) {
	t.Run("invalid URL", func(t *testing.T) {
		origURL := CreateCardURL
		CreateCardURL = "testInvalidURL%s"
		defer func() {
			CreateCardURL = origURL
		}()

		result, err := CreateCard("abc123a36ech8d75e160000f", "sample card unauthorized")
		assert.Equal(t, "", result)
		assert.NotNil(t, err.Error())
	})

	t.Run("unsupported protocol", func(t *testing.T) {
		origURL := CreateCardURL
		CreateCardURL = "testInvalidURL"
		defer func() {
			CreateCardURL = origURL
		}()

		result, err := CreateCard("abc123a36ech8d75e160000f", "sample card unauthorized")
		assert.Equal(t, "", result)
		assert.NotNil(t, err.Error())
	})

	t.Run("empty response", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()

		origURL := CreateCardURL
		CreateCardURL = ts.URL
		defer func() {
			CreateCardURL = origURL
		}()

		result, err := CreateCard("abc123a36ech8d75e160000f", "sample card unauthorized")
		assert.NotNil(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("error - unauthorized card creation", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		}))
		defer ts.Close()

		origURL := CreateCardURL
		CreateCardURL = ts.URL
		defer func() {
			CreateCardURL = origURL
		}()

		result, err := CreateCard("abc123a36ech8d75e160000f", "sample card unauthorized")
		assert.Equal(t, "", result)
		assert.Equal(t, "failed to create card - status code: 401", err.Error())
	})

	t.Run("successful", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"id": "abc123a36eaf8d78u160000f"}`))
			assert.Nil(t, err)
		}))
		defer ts.Close()

		origURL := CreateCardURL
		CreateCardURL = ts.URL
		defer func() {
			CreateCardURL = origURL
		}()

		result, err := CreateCard("abc123a36ech8d75e160000f", "sample card")
		assert.Equal(t, "abc123a36eaf8d78u160000f", result)
		assert.Nil(t, err)
	})
}
