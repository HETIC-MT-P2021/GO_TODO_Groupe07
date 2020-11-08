CREATE TABLE actions (
  action_id SERIAL PRIMARY KEY,

  content VARCHAR (255) NOT NULL,
  callback_at DATE,
  localisation VARCHAR (255),

  user_id VARCHAR (128) NOT NULL
);
