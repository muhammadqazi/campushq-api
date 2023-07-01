-- Description: Seed the database with some data

-- Create buildings
INSERT INTO buildings (name, code, description, number_of_rooms, created_at, updated_at)
VALUES
    ('Campus Tower', 'CT001', 'Modern office building with amenities', 20, NOW(), NOW()),
    ('Central Plaza', 'CP002', 'Prime location for retail businesses', 15, NOW(), NOW()),
    ('Skyline Apartments', 'SA003', 'Luxury residential complex with panoramic views', 10, NOW(), NOW()),
    ('Harmony Gardens', 'HG004', 'Beautiful community garden and event space', 12, NOW(), NOW()),
    ('Majestic Mall', 'MM005', 'Large shopping mall with diverse retail stores', 18, NOW(), NOW()),
    ('Prestige Tower', 'PT006', 'High-rise building with premium office spaces', 25, NOW(), NOW()),
    ('Serenity Residences', 'SR007', 'Tranquil living environment surrounded by nature', 30, NOW(), NOW()),
    ('Tech Hub', 'TH008', 'Innovation center for tech startups and entrepreneurs', 22, NOW(), NOW()),
    ('Lakeview Apartments', 'LA009', 'Lakefront residential complex with recreational facilities', 17, NOW(), NOW()),
    ('Grand Hotel', 'GH010', 'Elegant hotel offering luxurious accommodations', 14, NOW(), NOW());


-- Create rooms
INSERT INTO rooms (name, description, building_id, created_at, updated_at)
VALUES
    ('Room 101', 'Standard room with a view', 1, NOW(), NOW()),
    ('Room 102', 'Deluxe room with additional amenities', 1, NOW(), NOW()),
    ('Room 201', 'Spacious suite with a separate living area', 2, NOW(), NOW()),
    ('Room 202', 'Executive suite with luxury features', 2, NOW(), NOW()),
    ('Room 301', 'Cozy room perfect for solo travelers', 3, NOW(), NOW()),
    ('Room 302', 'Family room with multiple beds', 3, NOW(), NOW()),
    ('Room 401', 'Elegant room with a balcony', 4, NOW(), NOW()),
    ('Room 402', 'Penthouse suite with stunning city views', 4, NOW(), NOW()),
    ('Room 501', 'Accessible room for guests with disabilities', 5, NOW(), NOW()),
    ('Room 502', 'VIP suite with exclusive privileges', 5, NOW(), NOW());
