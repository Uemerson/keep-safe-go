package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

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

func decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("SECRET_KEY_32_BYTES")))
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func (ns *NoteService) AddNote(note *entity.NoteEntity) (*entity.NoteEntity, *exception.Exception) {
	errorsCauses := []exception.Causes{}
	if note.Text == "" {
		cause := exception.Causes{
			Message: "Missing param: text",
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
	if _, err := ns.nr.AddNote(note); err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	return note, nil
}

func (ns *NoteService) LoadNotes() ([]*entity.NoteEntity, *exception.Exception) {
	results, err := ns.nr.LoadNotes()
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	var notes []*entity.NoteEntity
	for _, r := range results {
		decoded, errDecode := base64.StdEncoding.DecodeString(r.Text)
		if errDecode != nil {
			return nil, exception.NewInternalServerError(errDecode.Error())
		}
		decrypted, errDecrypt := decrypt(decoded)
		if errDecrypt != nil {
			return nil, exception.NewInternalServerError(errDecrypt.Error())
		}
		notes = append(notes, &entity.NoteEntity{
			ID:        r.ID,
			Text:      string(decrypted),
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}
	return notes, nil
}

func (ns *NoteService) DeleteNote(id string) *exception.Exception {
	note, err := ns.nr.LoadNoteById(id)
	if err != nil {
		return exception.NewInternalServerError(err.Error())
	}
	if note == nil {
		return exception.NewNotFoundError(fmt.Sprintf("Note not found with this ID: %s", id))
	}
	if err := ns.nr.DeleteNote(id); err != nil {
		return exception.NewInternalServerError(err.Error())
	}
	return nil
}

func (ns *NoteService) LoadNoteById(id string) (*entity.NoteEntity, *exception.Exception) {
	note, err := ns.nr.LoadNoteById(id)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	if note == nil {
		return nil, exception.NewNotFoundError(fmt.Sprintf("Note not found with this ID: %s", id))
	}
	decoded, errDecode := base64.StdEncoding.DecodeString(note.Text)
	if errDecode != nil {
		return nil, exception.NewInternalServerError(errDecode.Error())
	}
	decrypted, errDecrypt := decrypt(decoded)
	if errDecrypt != nil {
		return nil, exception.NewInternalServerError(errDecrypt.Error())
	}
	note.Text = string(decrypted)
	return note, nil
}

func (ns *NoteService) SaveNote(id string, note *entity.NoteEntity) (*entity.NoteEntity, *exception.Exception) {
	n, err := ns.nr.LoadNoteById(id)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	if n == nil {
		return nil, exception.NewNotFoundError(fmt.Sprintf("Note not found with this ID: %s", id))
	}
	errorsCauses := []exception.Causes{}
	if note.Text == "" {
		cause := exception.Causes{
			Message: "Missing param: text",
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
	n.UpdatedAt = time.Now()
	n.Text = base64.StdEncoding.EncodeToString(encrypted)
	if err := ns.nr.SaveNote(id, n); err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	return n, nil
}
