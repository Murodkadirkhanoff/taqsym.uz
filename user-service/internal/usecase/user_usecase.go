package usecase

import (
	"errors"
	"time"

	"github.com/taqsym/taqsym.uz/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

// UserUseCase implements the domain.UserUseCase interface
type UserUseCase struct {
	repo domain.UserRepository
}

// NewUserUseCase creates a new UserUseCase
func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

// GetByID retrieves a user by its ID
func (uc *UserUseCase) GetByID(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	user, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetByUsername retrieves a user by its username
func (uc *UserUseCase) GetByUsername(username string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	user, err := uc.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetByEmail retrieves a user by its email
func (uc *UserUseCase) GetByEmail(email string) (*domain.User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetAll retrieves all users
func (uc *UserUseCase) GetAll() ([]*domain.User, error) {
	return uc.repo.GetAll()
}

// Create creates a new user
func (uc *UserUseCase) Create(user *domain.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	// Check if username already exists
	existingUser, err := uc.repo.GetByUsername(user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Check if email already exists
	existingUser, err = uc.repo.GetByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Set default values
	if user.Role == "" {
		user.Role = string(domain.UserRoleUser)
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return uc.repo.Create(user)
}

// Update updates an existing user
func (uc *UserUseCase) Update(user *domain.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if user.ID == "" {
		return errors.New("user ID cannot be empty")
	}

	if user.Username == "" {
		return errors.New("username cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	// Check if the user exists
	existingUser, err := uc.repo.GetByID(user.ID)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	// Check if username already exists (if changed)
	if user.Username != existingUser.Username {
		userWithSameUsername, err := uc.repo.GetByUsername(user.Username)
		if err != nil {
			return err
		}
		if userWithSameUsername != nil {
			return errors.New("username already exists")
		}
	}

	// Check if email already exists (if changed)
	if user.Email != existingUser.Email {
		userWithSameEmail, err := uc.repo.GetByEmail(user.Email)
		if err != nil {
			return err
		}
		if userWithSameEmail != nil {
			return errors.New("email already exists")
		}
	}

	// Preserve password, created_by and created_at from the existing user
	user.Password = existingUser.Password
	user.CreatedAt = existingUser.CreatedAt

	// Update the updated_at timestamp
	user.UpdatedAt = time.Now()

	return uc.repo.Update(user)
}

// Delete deletes a user by its ID
func (uc *UserUseCase) Delete(id string) error {
	if id == "" {
		return errors.New("user ID cannot be empty")
	}

	return uc.repo.Delete(id)
}

// Authenticate authenticates a user with username and password
func (uc *UserUseCase) Authenticate(username, password string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	// Get user by username
	user, err := uc.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid username or password")
	}

	// Check if the user is active
	if !user.Active {
		return nil, errors.New("user account is inactive")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}
