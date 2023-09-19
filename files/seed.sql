CREATE TABLE IF NOT EXISTS users (
    uid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(55) unique  not null ,
    email VARCHAR(75) unique not null,
    password VARCHAR(155) not null,
    full_name VARCHAR(155) not null,
    user_age int8 default 0,
    phone_number VARCHAR(155),
    institution VARCHAR(155),
    address TEXT DEFAULT '',
    major VARCHAR(55),
    entry_year int2 default 0,
    linkedin_url VARCHAR(100),
    line_id VARCHAR(55),
    team_id uuid,
    is_profile_verified boolean default false,
    inserted_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);

CREATE TABLE IF NOT EXISTS users_docs (
    uid uuid PRIMARY KEY,
    student_card_filename VARCHAR(155) DEFAULT '',
    student_card_url text DEFAULT '',
    student_card_status INT DEFAULT 0,
    student_card_rejection text DEFAULT '',

    self_portrait_filename VARCHAR(155) DEFAULT '',
    self_portrait_url text DEFAULT '',
    self_portrait_status INT DEFAULT 0,
    self_portrait_rejection text DEFAULT '',

    twibbon_filename VARCHAR(155) DEFAULT '',
    twibbon_url text DEFAULT '',
    twibbon_status INT DEFAULT 0,
    twibbon_rejection text DEFAULT '',

    enrollment_filename VARCHAR(155) DEFAULT '',
    enrollment_url text DEFAULT '',
    enrollment_status INT DEFAULT 0,
    enrollment_rejection text DEFAULT '',

    is_doc_verified boolean default false
);

CREATE TABLE IF NOT EXISTS teams(
    team_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    team_name VARCHAR(55) UNIQUE NOT NULL,
    team_leader_id UUID unique not null ,
    team_member_mails TEXT[],
    is_active boolean default true,
    submission_1_filename VARCHAR(155) DEFAULT '',
    submission_2_filename VARCHAR(155) DEFAULT ''
);

create TABLE IF NOT EXISTS teams_docs(
    team_id uuid PRIMARY KEY,
    payment_filename  VARCHAR(155) DEFAULT '',
    payment_url text DEFAULT '',
    payment_status INT DEFAULT 0,
    payment_rejection text DEFAULT '',
    submission_1_filename   varchar(155) default ''::STRING,
    submission_1_url        text         default ''::STRING,
    submission_1_lastupdate timestamp,
    submission_2_filename   varchar(155) default ''::STRING,
    submission_2_url        text         default ''::STRING,
    submission_2_lastupdate timestamp
);

CREATE TABLE IF NOT EXISTS teams_code(
    team_id uuid PRIMARY KEY,
    code VARCHAR(55) UNIQUE,
    team_member_mails TEXT[],
    used int8 default 2
);

CREATE TABLE IF NOT EXISTS admins (
    uid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    username varchar(55) unique  not null ,
    password VARCHAR(155) not null,
    full_name VARCHAR(155) not null,
    inserted_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp
);