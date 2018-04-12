package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/youkoulayley/api-collection/models"
	"github.com/youkoulayley/api-collection/repositories"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

