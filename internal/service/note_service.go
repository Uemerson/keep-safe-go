package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"

	"github.com/Uemerson/keep-safe-go/internal/entity"
	"github.com/Uemerson/keep-safe-go/internal/exception"
	"github.com/Uemerson/keep-safe-go/internal/repository"
)

type NoteService struct {
	nr *repository.NoteRepository
}

func NewNoteService(nr *repository.NoteRepository) *NoteService {
	return &NoteService{nr: nr}
}

func encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("SECRET_KEY_32_BYTES")))
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

func (ns *NoteService) CreateNote(note *entity.NoteEntity) (*entity.NoteEntity, *exception.Exception) {
	errorsCauses := []exception.Causes{}
	if note.Text == "" {
		cause := exception.Causes{
			Message: "missing param text",
			Field:   "text",
		}
		errorsCauses = append(errorsCauses, cause)
	}
	if len(errorsCauses) > 0 {
		return nil, exception.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	}
	encrypted, errEncrypt := encrypt([]byte(note.Text))
	if errEncrypt != nil {
		return nil, exception.NewInternalServerError(errEncrypt.Error())
	}
	note.Text = base64.StdEncoding.EncodeToString(encrypted)
	if _, err := ns.nr.SaveNote(note); err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	return note, nil
}
