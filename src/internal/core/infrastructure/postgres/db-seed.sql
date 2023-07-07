-- Insert records for table `buildings`
INSERT INTO buildings (name, code, description, number_of_rooms)
VALUES
   ('Building Alpha', 'ALP001', 'Description for Building Alpha', 10),
   ('Building Bravo', 'BRV002', 'Description for Building Bravo', 8),
   ('Building Charlie', 'CHL003', 'Description for Building Charlie', 12),
   ('Building Delta', 'DLT004', 'Description for Building Delta', 6),
   ('Building Echo', 'ECH005', 'Description for Building Echo', 15),
   ('Building Foxtrot', 'FOX006', 'Description for Building Foxtrot', 10),
   ('Building Golf', 'GLF007', 'Description for Building Golf', 8),
   ('Building Hotel', 'HTL008', 'Description for Building Hotel', 12),
   ('Building India', 'IND009', 'Description for Building India', 6),
   ('Building Juliet', 'JLT010', 'Description for Building Juliet', 15);

-- Insert records for table `rooms`
INSERT INTO rooms (name, description, capacity, building_id)
VALUES
   ('Room 101', 'Description for Room 101', 20, 1),
   ('Room 102', 'Description for Room 102', 15, 1),
   ('Room 201', 'Description for Room 201', 10, 2),
   ('Room 202', 'Description for Room 202', 12, 2),
   ('Room 301', 'Description for Room 301', 18, 3),
   ('Room 302', 'Description for Room 302', 14, 3),
   ('Room 401', 'Description for Room 401', 16, 4),
   ('Room 402', 'Description for Room 402', 20, 4),
   ('Room 501', 'Description for Room 501', 22, 5),
   ('Room 502', 'Description for Room 502', 25, 5);

-- Insert records for table `staffs`
INSERT INTO staffs (staff_id, username, first_name, surname, sex, role, salary, status, office_id)
VALUES
   (1, 'john_doe', 'John', 'Doe', 'Male', 'Teacher', '5000', 'Active', 1),
   (2, 'jane_smith', 'Jane', 'Smith', 'Female', 'Teacher', '4500', 'Active', 2),
   (3, 'robert_johnson', 'Robert', 'Johnson', 'Male', 'Staff', '4000', 'Active', 3),
   (4, 'emily_wilson', 'Emily', 'Wilson', 'Female', 'Staff', '3500', 'Active', 4),
   (5, 'michael_brown', 'Michael', 'Brown', 'Male', 'Teacher', '5000', 'Active', 5),
   (6, 'sophia_miller', 'Sophia', 'Miller', 'Female', 'Teacher', '4500', 'Active', 6),
   (7, 'david_anderson', 'David', 'Anderson', 'Male', 'Staff', '4000', 'Active', 7),
   (8, 'olivia_clark', 'Olivia', 'Clark', 'Female', 'Staff', '3500', 'Active', 8),
   (9, 'william_thompson', 'William', 'Thompson', 'Male', 'Teacher', '5000', 'Active', 9),
   (10, 'ava_hall', 'Ava', 'Hall', 'Female', 'Teacher', '4500', 'Active', 10);

-- Insert records for table `faculties`
INSERT INTO faculties (name, description, faculty_email, faculty_phone_number, dean_id, vice_dean_id)
VALUES
   ('Faculty of Science', 'Description for Faculty of Science', 'science@university.edu', '+1234567890', 1, 2),
   ('Faculty of Arts', 'Description for Faculty of Arts', 'arts@university.edu', '+0987654321', 3, 4),
   ('Faculty of Engineering', 'Description for Faculty of Engineering', 'engineering@university.edu', '+9876543210', 5, 6),
   ('Faculty of Business', 'Description for Faculty of Business', 'business@university.edu', '+0123456789', 7, 8),
   ('Faculty of Medicine', 'Description for Faculty of Medicine', 'medicine@university.edu', '+2468135790', 9, 10);

-- Insert records for table `departments`
INSERT INTO departments (name, code, description, number_of_semesters, department_email, department_phone_number, tution_per_semester, head_of_department_id, faculty_id)
VALUES
   ('Computer Science', 'CSC', 'Description for Computer Science Department', 8, 'csc@university.edu', '+1122334455', 5000, 1, 1),
   ('English Literature', 'ENG', 'Description for English Literature Department', 6, 'eng@university.edu', '+5566778899', 4500, 2, 2),
   ('Mechanical Engineering', 'MEC', 'Description for Mechanical Engineering Department', 8, 'mec@university.edu', '+9900112233', 5500, 3, 3),
   ('Business Administration', 'BUS', 'Description for Business Administration Department', 6, 'bus@university.edu', '+4433221100', 4800, 4, 4),
   ('Medicine', 'MED', 'Description for Medicine Department', 12, 'med@university.edu', '+7788990011', 6000, 5, 5);

-- Insert records for table `students`
INSERT INTO students (student_id, first_name, surname, sex, role, status, access_status, acceptance_type, semester, graduation_date, supervisor_id, department_id, created_at, updated_at, is_active)
VALUES
   (21906778, 'John', 'Doe', 'Male', 'Student', 'Active', 'Allowed', 'Accepted', 'Fall 2022', NULL, 1, 1, NOW(), NOW(), true),
   (21906779, 'Jane', 'Smith', 'Female', 'Student', 'Active', 'Allowed', 'Accepted', 'Spring 2023', NULL, 2, 2, NOW(), NOW(), true),
   (21906710, 'Robert', 'Johnson', 'Male', 'Student', 'Active', 'Allowed', 'Accepted', 'Fall 2021', NULL, 3, 3, NOW(), NOW(), true),
   (21906711, 'Emily', 'Wilson', 'Female', 'Student', 'Active', 'Allowed', 'Accepted', 'Spring 2024', NULL, 4, 4, NOW(), NOW(), true),
   (21906712, 'Michael', 'Brown', 'Male', 'Student', 'Active', 'Allowed', 'Accepted', 'Fall 2022', NULL, 5, 5, NOW(), NOW(), true),
   (21906713, 'Sophia', 'Miller', 'Female', 'Student', 'Active', 'Allowed', 'Accepted', 'Spring 2023', NULL, 6, 1, NOW(), NOW(), true),
   (21906714, 'David', 'Anderson', 'Male', 'Student', 'Active', 'Allowed', 'Accepted', 'Fall 2021', NULL, 7, 2, NOW(), NOW(), true),
   (21906715, 'Olivia', 'Clark', 'Female', 'Student', 'Active', 'Allowed', 'Accepted', 'Spring 2024', NULL, 8, 3, NOW(), NOW(), true),
   (21906716, 'William', 'Thompson', 'Male', 'Student', 'Active', 'Allowed', 'Accepted', 'Fall 2022', NULL, 9, 4, NOW(), NOW(), true),
   (21906717, 'Ava', 'Hall', 'Female', 'Student', 'Active', 'Allowed', 'Accepted', 'Spring 2023', NULL, 10, 5, NOW(), NOW(), true);

-- Insert records for table `student_accounts`
INSERT INTO student_accounts (scholarship_percentage, discount_percentage, discount_type, total_debt, student_id)
VALUES
   (10, 5, 'Percentage', 5000, 21906778),
   (15, 7, 'Percentage', 6000,  21906779),
   (20, 10, 'Percentage', 7000, 21906710),
   (25, 12, 'Percentage', 8000, 21906711),
   (30, 15, 'Percentage', 9000, 21906712),
   (35, 17, 'Percentage', 10000, 21906713),
   (40, 20, 'Percentage', 11000, 21906714),
   (45, 22, 'Percentage', 12000, 21906715),
   (50, 25, 'Percentage', 13000, 21906716),
   (55, 27, 'Percentage', 14000, 21906717);

-- Insert records for table `student_invoices`
INSERT INTO student_invoices (invoice_number, invoice_date, invoice_due_date, invoice_amount, invoice_status, invoice_type, invoice_description, term, account_id)
VALUES
   ('INV001', NOW(), NOW() + INTERVAL '30 days', 2000, 'Pending', 'Tuition', 'Tuition fees for the semester', 'Spring 2023', 1),
   ('INV002', NOW(), NOW() + INTERVAL '30 days', 1500, 'Paid', 'Tuition', 'Tuition fees for the semester', 'Fall 2022', 2),
   ('INV003', NOW(), NOW() + INTERVAL '30 days', 1800, 'Pending', 'Tuition', 'Tuition fees for the semester', 'Spring 2023', 3),
   ('INV004', NOW(), NOW() + INTERVAL '30 days', 2200, 'Pending', 'Tuition', 'Tuition fees for the semester', 'Fall 2022', 4),
   ('INV005', NOW(), NOW() + INTERVAL '30 days', 1600, 'Paid', 'Tuition', 'Tuition fees for the semester', 'Spring 2023', 5),
   ('INV006', NOW(), NOW() + INTERVAL '30 days', 1900, 'Paid', 'Tuition', 'Tuition fees for the semester', 'Fall 2022', 6),
   ('INV007', NOW(), NOW() + INTERVAL '30 days', 2300, 'Pending', 'Tuition', 'Tuition fees for the semester', 'Spring 2023', 7),
   ('INV008', NOW(), NOW() + INTERVAL '30 days', 1700, 'Pending', 'Tuition', 'Tuition fees for the semester', 'Fall 2022', 8),
   ('INV009', NOW(), NOW() + INTERVAL '30 days', 2000, 'Paid', 'Tuition', 'Tuition fees for the semester', 'Spring 2023', 9),
   ('INV010', NOW(), NOW() + INTERVAL '30 days', 2400, 'Pending', 'Tuition', 'Tuition fees for the semester', 'Fall 2022', 10);

-- Insert records for table `student_payments`
INSERT INTO student_payments (payment_date, payment_amount, payment_method, payment_description, payment_status, account_id)
VALUES
   (NOW(), 1000, 'Credit Card', 'Payment towards tuition fees', 'Paid', 1),
   (NOW(), 800, 'Bank Transfer', 'Payment towards tuition fees', 'Paid', 2),
   (NOW(), 1200, 'Cash', 'Payment towards tuition fees', 'Pending', 3),
   (NOW(), 1500, 'Credit Card', 'Payment towards tuition fees', 'Paid', 4),
   (NOW(), 900, 'Bank Transfer', 'Payment towards tuition fees', 'Pending', 5),
   (NOW(), 1100, 'Cash', 'Payment towards tuition fees', 'Paid', 6),
   (NOW(), 1300, 'Credit Card', 'Payment towards tuition fees', 'Pending', 7),
   (NOW(), 950, 'Bank Transfer', 'Payment towards tuition fees', 'Paid', 8),
   (NOW(), 1050, 'Cash', 'Payment towards tuition fees', 'Paid', 9),
   (NOW(), 1400, 'Credit Card', 'Payment towards tuition fees', 'Pending', 10);

-- Insert records for table `personal_information`
INSERT INTO personal_information (nationality, city, address, dob, place_of_birth, local_address, guardian_phone_number, local_phone_number, father_name, mother_name, passport_number, passport_issue_date, passport_expiry_date, passport_issue_place, passport_issuing_authority, id_card_number, student_id, staff_id, is_staff)
VALUES
   ('USA', 'New York', '123 ABC Street', '1995-06-15', 'New York', '456 XYZ Street', '1234567890', '9876543210', 'John Doe', 'Jane Doe', 'A1234567', '2022-01-01', '2027-01-01', 'New York', 'US Government', '1234567890', 21906778, NULL, false),
    ('USA', 'Los Angeles', '456 XYZ Street', '1996-03-12', 'Los Angeles', '789 PQR Street', '9876543210', '1234567890', 'Michael Smith', 'Emily Smith', 'B2345678', '2022-02-01', '2027-02-01', 'Los Angeles', 'US Government', '9876543210', 21906779, NULL, false),
    ('USA', 'Chicago', '789 PQR Street', '1997-09-18', 'Chicago', '123 ABC Street', '1234567890', '9876543210', 'Robert Johnson', 'Emma Johnson', 'C9876543', '2022-03-01', '2027-03-01', 'Chicago', 'US Government', '1234567890', 21906710, NULL, false),
    ('USA', 'Houston', '321 DEF Street', '1998-12-05', 'Houston', '654 LMN Street', '9876543210', '1234567890', 'William Davis', 'Olivia Davis', 'D8765432', '2022-04-01', '2027-04-01', 'Houston', 'US Government', '9876543210', 21906711, NULL, false),
    ('USA', 'San Francisco', '654 LMN Street', '1999-07-22', 'San Francisco', '987 QWE Street', '1234567890', '9876543210', 'David Anderson', 'Sophia Anderson', 'E7654321', '2022-05-01', '2027-05-01', 'San Francisco', 'US Government', '1234567890', 21906712, NULL, false),
    ('USA', 'New York', '123 ABC Street', '1995-06-15', 'New York', '456 XYZ Street', '1234567890', '9876543210', 'John Doe', 'Jane Doe', 'A1234567', '2022-01-01', '2027-01-01', 'New York', 'US Government', '1234567890', 21906713, NULL, false),
    ('USA', 'Los Angeles', '456 XYZ Street', '1996-03-12', 'Los Angeles', '789 PQR Street', '9876543210', '1234567890', 'Michael Smith', 'Emily Smith', 'B2345678', '2022-02-01', '2027-02-01', 'Los Angeles', 'US Government', '9876543210', 21906714, NULL, false),
    ('USA', 'Chicago', '789 PQR Street', '1997-09-18', 'Chicago', '123 ABC Street', '1234567890', '9876543210', 'Robert Johnson', 'Emma Johnson', 'C9876543', '2022-03-01', '2027-03-01', 'Chicago', 'US Government', '1234567890', 21906715, NULL, false),
    ('USA', 'Houston', '321 DEF Street', '1998-12-05', 'Houston', '654 LMN Street', '9876543210', '1234567890', 'William Davis', 'Olivia Davis', 'D8765432', '2022-04-01', '2027-04-01', 'Houston', 'US Government', '9876543210', 21906716, NULL, false),
    ('USA', 'San Francisco', '654 LMN Street', '1999-07-22', 'San Francisco', '987 QWE Street', '1234567890', '9876543210', 'David Anderson', 'Sophia Anderson', 'E7654321', '2022-05-01', '2027-05-01', 'San Francisco', 'US Government', '1234567890', 21906717, NULL, false);
