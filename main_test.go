package main

// func TestReqCreateUser(t *testing.T) {
// 	factoryRepo, err := createRepositoryFactoryInmem()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	userRepo, err := createRepositoryUserInmem(factoryRepo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	username := "random_username_name"
// 	body := []byte(fmt.Sprintf(`{"username":"%s"}`, username))

// 	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rec := httptest.NewRecorder()
// 	handler := http.HandlerFunc(reqCreateUser(userRepo, factoryRepo))
// 	handler.ServeHTTP(rec, req)

// 	if rec.Code != http.StatusOK {
// 		t.Fatalf("wrong status code: got %v want %v", rec.Code, http.StatusOK)
// 	}

// 	if rec.Body.String() != "true" {
// 		t.Fatal("unexpected body")
// 	}

// 	user, err := userRepo.GetUserByUsername(username)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if len(user.GetUsername()) == 0 {
// 		t.Fatal("user not created")
// 	}

// 	if user.GetUsername() != username {
// 		t.Fatal("wrong username")
// 	}
// }

// func TestReqGetDashboard(t *testing.T) {
// 	factoryRepo, err := createRepositoryFactoryInmem()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	userRepo, err := createRepositoryUserInmem(factoryRepo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	username := "random_username_name"
// 	_, err = createUser(username, userRepo, factoryRepo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	req, err := http.NewRequest("GET", fmt.Sprintf("/dashboard?username=%s", username), nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rec := httptest.NewRecorder()
// 	handler := http.HandlerFunc(reqGetDashboard(userRepo, factoryRepo))
// 	handler.ServeHTTP(rec, req)

// 	if rec.Code != http.StatusOK {
// 		t.Fatalf("wrong status code: got %v want %v", rec.Code, http.StatusOK)
// 	}

// 	// TODO: should test the body
// 	// if rec.Body.String() != "true" {
// 	// 	t.Fatal("unexpected body")
// 	// }
// }
