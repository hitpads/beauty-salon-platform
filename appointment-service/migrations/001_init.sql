CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS services (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT,
    price_cents INT NOT NULL,
    duration_minutes INT NOT NULL
);

CREATE TABLE IF NOT EXISTS appointments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    master_id UUID NOT NULL,
    service_id UUID NOT NULL,
    start_time TIMESTAMP NOT NULL,
    status TEXT NOT NULL DEFAULT 'scheduled'
    -- FOREIGN KEYs можно добавить, если user/master таблицы есть
);
