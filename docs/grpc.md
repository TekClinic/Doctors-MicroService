## gRPC Functions

### GetDoctor

Retrieves the details of a specific doctor by their ID.

**Request:**

```protobuf
message GetDoctorRequest {
  string token = 1; // Authentication token
  int32 id = 2; // ID of the doctor
}
```

**Response:**

```protobuf
message GetDoctorResponse {
  Doctor doctor = 1; // Details of the doctor
}
```

**Errors:**

- `Unauthenticated` - Token is not valid or expired.
- `PermissionDenied` - Token is not authorized with the *admin* role.
- `NotFound` - Doctor with the given ID does not exist.

---

### GetDoctorsIDs

Retrieves a list of doctor IDs with pagination support.

**Request:**

```protobuf
message GetDoctorsIDsRequest {
  string token = 1; // Authentication token
  int32 limit = 2; // Maximum number of results to return
  int32 offset = 3; // Offset for pagination
  string search = 4; // Search term for filtering results (optional)
}
```

**Response:**

```protobuf
message GetDoctorsIDsResponse {
  int32 count = 1; // Total number of doctors
  repeated int32 results = 2; // List of doctor IDs
}
```

**Errors:**

- `Unauthenticated` - Token is not valid or expired.
- `PermissionDenied` - Token is not authorized with the *admin* role.
- `InvalidArgument` - `offset`, `limit`, or `search` parameters are invalid.

---

### CreateDoctor

Creates a new doctor record with the provided details.

**Request:**

```protobuf
message CreateDoctorRequest {
  string token = 1; // Authentication token
  string name = 2; // Name of the doctor
  Doctor.Gender gender = 3; // Gender of the doctor (optional)
  string phone_number = 4; // Phone number of the doctor
  repeated string specialities = 5; // Specialties of the doctor
  string special_note = 6; // Special notes regarding the doctor (optional)
}
```

**Response:**

```protobuf
message CreateDoctorResponse {
  int32 id = 1; // ID of the newly created doctor
}
```

**Errors:**

- `Unauthenticated` - Token is not valid or expired.
- `PermissionDenied` - Token is not authorized with the *admin* role.
- `InvalidArgument` - Required doctor information is missing or malformed.

---

### DeleteDoctor

Deletes a doctor record by their ID.

**Request:**

```protobuf
message DeleteDoctorRequest {
  string token = 1; // Authentication token
  int32 id = 2; // ID of the doctor to be deleted
}
```

**Response:**

```protobuf
message DeleteDoctorResponse {}
```

**Errors:**

- `Unauthenticated` - Token is not valid or expired.
- `PermissionDenied` - Token is not authorized with the *admin* role.
- `NotFound` - Doctor with the given ID does not exist.

---

### UpdateDoctor

Updates the details of an existing doctor.

**Request:**

```protobuf
message UpdateDoctorRequest {
  string token = 1; // Authentication token
  Doctor doctor = 2; // Updated doctor details
}
```

**Response:**

```protobuf
message UpdateDoctorResponse {
  int32 id = 1; // ID of the updated doctor
}
```

**Errors:**

- `Unauthenticated` - Token is not valid or expired.
- `PermissionDenied` - Token is not authorized with the *admin* role.
- `InvalidArgument` - Updated doctor information is missing or malformed.
- `NotFound` - Doctor with the given ID does not exist.

---

## Model Definition

```protobuf
message Doctor {
  enum Gender {
    UNSPECIFIED = 0;
    MALE = 1;
    FEMALE = 2;
  }

  int32 id = 1; // ID of the doctor
  bool active = 2; // Flag indicating if the doctor is active
  string name = 3; // Name of the doctor
  Gender gender = 4; // Gender of the doctor
  string phone_number = 5; // Phone number of the doctor
  repeated string specialities = 6; // Specialties of the doctor
  string special_note = 7; // Special notes regarding the doctor
}
```