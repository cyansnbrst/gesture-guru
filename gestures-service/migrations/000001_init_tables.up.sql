CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE gestures (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    video_url TEXT,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Категории для РЖЯ
INSERT INTO categories (name) VALUES 
('Азбука (Дактиль)'),
('Базовые слова'),
('Числа');

-- Жесты РЖЯ
INSERT INTO gestures (name, description, video_url, category_id) VALUES
('А (дактиль)', 'Указательный палец вверх, остальные собраны в кулак', 'https://storage.rjya/azbuka/a.mp4', 1),
('М (дактиль)', 'Три пальца (большой, указательный, средний) соединены', 'https://storage.rjya/azbuka/m.mp4', 1),
('Привет', 'Поднять раскрытую ладонь ко лбу и отвести вперед', 'https://storage.rjya/basic/privet.mp4', 2),
('Спасибо', 'Кисть у губ, движение от себя с наклоном головы', 'https://storage.rjya/basic/spasibo.mp4', 2),
('5 (число)', 'Раскрытая ладонь с растопыренными пальцами', 'https://storage.rjya/numbers/5.mp4', 3);