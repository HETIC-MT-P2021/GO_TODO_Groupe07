CREATE TABLE actions (
  action_id SERIAL PRIMARY KEY,

  content VARCHAR (255) NOT NULL,
  callback_at DATE,
  localisation VARCHAR (255),

  user_id INTEGER references users(user_id)
);

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,

    name VARCHAR (64) NOT NULL,
    discord_tag INTEGER NOT NULL,
    discord_server_id BIGINT NOT NULL,
);
