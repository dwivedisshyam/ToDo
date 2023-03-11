CREATE TABLE public."user" (
    id serial4 NOT NULL,
    username varchar(30) NOT NULL,
    "password" varchar(255) NOT NULL,
    full_name varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    "role" varchar(10) NOT NULL DEFAULT 'User'::character varying,
    created_at timestamptz NOT NULL,
    CONSTRAINT "users_pk" PRIMARY KEY (id),
    CONSTRAINT "users_username_key" UNIQUE (username)
);

CREATE TABLE public."task" (
    id serial4 NOT NULL,
    user_id int8 NOT NULL,
    title varchar(50) NOT NULL,
    description varchar(255) NULL,
    due_date timestamptz NULL,
    created_at timestamptz NOT NULL,
    CONSTRAINT "tasks_pk" PRIMARY KEY (id),
    CONSTRAINT "tasks_fk0" FOREIGN KEY (user_id) REFERENCES public."user"(id)
);