BEGIN TRANSACTION;

CREATE TABLE users (
	  id SERIAL PRIMARY KEY,
	  username VARCHAR(100),
	  password VARCHAR(1000),
	  email VARCHAR(100)
);
	
INSERT INTO users (id, username, password, email) VALUES (1, 'suraj', 'suraj', 'sapatil@live.com');

CREATE TABLE categorys (
	  id SERIAL PRIMARY KEY,
	  name VARCHAR(1000) NOT NULL,
	  user_id INTEGER REFERENCES users(id)
);

INSERT INTO categorys (id, name, user_id) VALUES (1, 'TaskApp', 1);

CREATE TABLE status (
	  id SERIAL PRIMARY KEY,
	  status VARCHAR(50) NOT NULL
);

INSERT INTO status (id, status) VALUES 
(1, 'COMPLETE'),
(2, 'PENDING'),
(3, 'DELETED');

CREATE TABLE task (
	  id SERIAL PRIMARY KEY,
	  title VARCHAR(100),
	  content TEXT,
	  created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  last_modified_at TIMESTAMP,
	  finish_date TIMESTAMP,
	  priority INTEGER,
	  cat_id INTEGER REFERENCES categorys(id),
	  task_status_id INTEGER REFERENCES status(id),
	  due_date TIMESTAMP,
	  user_id INTEGER REFERENCES users(id),
	  hide INTEGER
);

INSERT INTO task (id, title, content, created_date, last_modified_at, finish_date, priority, cat_id, task_status_id, due_date, user_id, hide)
	VALUES 
	(1, 'Publishing on GitHub', 'Publish the source of tasks...', '2015-11-12 15:30:59', '2015-11-21 14:19:22', '2015-11-17 17:02:18', 3, 1, 1, NULL, 1, 0);

	CREATE TABLE comments (
		  id SERIAL PRIMARY KEY,
		  content TEXT,
		  taskID INTEGER REFERENCES task(id),
		  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		  user_id INTEGER REFERENCES users(id)
	);

	CREATE TABLE files (
		  name VARCHAR(1000) NOT NULL,
		  autoName VARCHAR(255) NOT NULL,
		  user_id INTEGER REFERENCES users(id),
		  created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	COMMIT;
