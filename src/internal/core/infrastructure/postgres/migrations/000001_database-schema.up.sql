-- Database initialization script.


--
-- Table structure for table `buildings`
--

CREATE TABLE buildings (
   building_id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   code VARCHAR(50) NOT NULL UNIQUE,
   description VARCHAR(255) NOT NULL,
   number_of_rooms INT NOT NULL,

   created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMPTZ NULL,
   deleted_at TIMESTAMPTZ,
   is_active BOOLEAN NOT NULL DEFAULT TRUE
);

--
-- Table structure for table `rooms`
--

CREATE TABLE rooms (
   room_id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   description VARCHAR(255) NOT NULL,
   capacity INT NULL,
   
   created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMPTZ NULL,
   deleted_at TIMESTAMPTZ,
   is_active BOOLEAN NOT NULL DEFAULT TRUE,

   building_id INT NOT NULL,

   CONSTRAINT fk_room_building_id FOREIGN KEY (building_id) REFERENCES buildings(building_id)
);

--
-- Table structure for table `staffs`
--

CREATE TABLE staffs (
   staff_id INT NOT NULL PRIMARY KEY,
   username VARCHAR(50) NULL UNIQUE,
   first_name VARCHAR(50) NOT NULL,
   surname VARCHAR(50) NOT NULL,
   sex VARCHAR(50) NOT NULL,
   role VARCHAR(50) NOT NULL,
   salary VARCHAR(50) NOT NULL,
   status VARCHAR(50) NOT NULL,
   office_id INT NOT NULL,

   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP,

   CONSTRAINT fk_staff_office_id FOREIGN KEY (office_id) REFERENCES rooms(room_id)
);

--
-- Table structure for table `faculties`
--

CREATE TABLE faculties (
   faculty_id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   description VARCHAR(255) NOT NULL,
   faculty_email VARCHAR(50) NOT NULL,
   faculty_phone_number VARCHAR(50) NOT NULL,

   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP,

   dean_id INT NOT NULL,
   vice_dean_id INT NOT NULL,

   CONSTRAINT fk_faculty_dean_id FOREIGN KEY (dean_id) REFERENCES staffs(staff_id),
   CONSTRAINT fk_faculty_vice_dean_id FOREIGN KEY (vice_dean_id) REFERENCES staffs(staff_id)
);

--
-- Table structure for table `staffs`
--

CREATE TABLE departments (
   department_id SERIAL PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   code VARCHAR(50) NOT NULL UNIQUE,
   description VARCHAR(255) NOT NULL,
   number_of_semesters INT NOT NULL,
   department_email VARCHAR(50) NOT NULL,
   department_phone_number VARCHAR(50) NOT NULL,
   tution_per_semester INT NOT NULL,

   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP,

   head_of_department_id INT NOT NULL,
   faculty_id INT NOT NULL,

   CONSTRAINT fk_department_head_of_department_id FOREIGN KEY (head_of_department_id) REFERENCES staffs(staff_id),
   CONSTRAINT fk_department_faculty_id FOREIGN KEY (faculty_id) REFERENCES faculties(faculty_id)
);

--
-- Table structure for table `students`
--

CREATE TABLE students (
    student_id INT NOT NULL PRIMARY KEY,
    username VARCHAR(50) NULL UNIQUE,
    first_name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    sex VARCHAR(50) NOT NULL,
    role VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    access_status VARCHAR(50) NOT NULL,
    acceptance_type VARCHAR(50) NOT NULL,
    semester VARCHAR(50) NOT NULL,
    graduation_date TIMESTAMP,
    supervisor_id INT,
    department_id INT,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    is_active BOOLEAN,

   CONSTRAINT fk_student_supervisor_id FOREIGN KEY(supervisor_id) REFERENCES staffs(staff_id),
   CONSTRAINT fk_student_department_id FOREIGN KEY(department_id) REFERENCES departments(department_id)
);

--
-- Table structure for table `accounts`
--

CREATE TABLE student_accounts (
      account_id SERIAL PRIMARY KEY,
      scholarship_percentage INT NOT NULL,
      discount_percentage INT NOT NULL,
      discount_type VARCHAR(50) NOT NULL,
      total_debt INT NOT NULL,
   
      updated_at TIMESTAMP NULL,
   
      student_id INT NULL,
   
      CONSTRAINT fk_account_student_id FOREIGN KEY(student_id) REFERENCES students(student_id)
);

--
-- Table structure for table `invoices`
--

CREATE TABLE student_invoices (
      invoice_id SERIAL PRIMARY KEY,
      invoice_number VARCHAR(50) NOT NULL,
      invoice_date TIMESTAMP NOT NULL,
      invoice_due_date TIMESTAMP NOT NULL,
      invoice_amount INT NOT NULL,
      invoice_status VARCHAR(50) NOT NULL,
      invoice_type VARCHAR(50) NOT NULL,
      invoice_description VARCHAR(255) NOT NULL,
      term VARCHAR(50) NOT NULL,

      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP NULL,

      account_id INT NOT NULL,

      CONSTRAINT fk_invoice_account_id FOREIGN KEY(account_id) REFERENCES student_accounts(account_id)
);

--
-- Table structure for table `payments`
--

CREATE TABLE student_payments (
      payment_id SERIAL PRIMARY KEY,
      payment_date TIMESTAMP NOT NULL,
      payment_amount INT NOT NULL,
      payment_method VARCHAR(50) NOT NULL,
      payment_description VARCHAR(255) NOT NULL,
      payment_status VARCHAR(50) NOT NULL,

      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP NULL,

      account_id INT NOT NULL,

      CONSTRAINT fk_payment_account_id FOREIGN KEY(account_id) REFERENCES student_accounts(account_id)
);

--
-- Table structure for table `personal_information`
--

CREATE TABLE personal_information (
    pid SERIAL PRIMARY KEY,
    nationality VARCHAR(50) NOT NULL,
    city VARCHAR(50) NULL,
    address VARCHAR(50) NULL,
    dob DATE NOT NULL,
    place_of_birth VARCHAR(50) NOT NULL,
    local_address VARCHAR(50) NULL,
    guardian_phone_number VARCHAR(50) NOT NULL,
    local_phone_number VARCHAR(50) NULL,
    father_name VARCHAR(50) NOT NULL,
    mother_name VARCHAR(50) NOT NULL,

    passport_number VARCHAR(50) NOT NULL,
    passport_issue_date DATE NOT NULL,
    passport_expiry_date DATE NOT NULL,
    passport_issue_place VARCHAR(50) NULL,
    passport_issuing_authority VARCHAR(50) NOT NULL,
    id_card_number VARCHAR(50) NULL,

    student_id INT NULL,
    staff_id INT NULL,
    is_staff BOOLEAN NOT NULL DEFAULT FALSE,

    updated_at TIMESTAMP NULL,

    CONSTRAINT fk_personal_information_student_id FOREIGN KEY(student_id) REFERENCES students(student_id),
    CONSTRAINT fk_personal_information_staff_id FOREIGN KEY(staff_id) REFERENCES staffs(staff_id)
);

