-- +goose Up
-- +goose StatementBegin

-- * reuse users table from original schema
CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    oid         VARCHAR(64)  NOT NULL UNIQUE,
    firstname   VARCHAR(255) NOT NULL,
    lastname    VARCHAR(255) NOT NULL,
    email       VARCHAR(255) UNIQUE,
    picture_url VARCHAR(255) NULL,
    is_admin    BOOLEAN      NOT NULL,
    metadata    JSONB        NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE topics
(
    id           BIGSERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    label        VARCHAR(255) NOT NULL,
    embedding_no BIGINT       NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE courses
(
    id                BIGSERIAL PRIMARY KEY,
    name              VARCHAR(255) NOT NULL,
    description       TEXT         NULL,
    prompt_instruction TEXT         NULL,
    created_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE course_managers
(
    id         BIGSERIAL PRIMARY KEY,
    course_id  BIGINT REFERENCES courses (id) ON DELETE CASCADE NOT NULL,
    user_id    BIGINT REFERENCES users (id) ON DELETE CASCADE   NOT NULL,
    created_at TIMESTAMP                                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (course_id, user_id)
);

CREATE TABLE enrolls
(
    id         BIGSERIAL PRIMARY KEY,
    course_id  BIGINT REFERENCES courses (id) ON DELETE CASCADE NOT NULL,
    user_id    BIGINT REFERENCES users (id) ON DELETE CASCADE   NOT NULL,
    created_at TIMESTAMP                                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (course_id, user_id)
);

CREATE TABLE contents
(
    id         BIGSERIAL PRIMARY KEY,
    enroll_id  BIGINT REFERENCES enrolls (id) ON DELETE CASCADE NOT NULL,
    title      VARCHAR(255)                                          NOT NULL,
    created_at TIMESTAMP                                             NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                                             NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE content_sections
(
    id          BIGSERIAL PRIMARY KEY,
    content_id  BIGINT REFERENCES contents (id) ON DELETE CASCADE NOT NULL,
    section_no  INTEGER                                              NOT NULL,
    title       VARCHAR(255)                                         NOT NULL,
    subtitle    VARCHAR(255)                                         NULL,
    content     TEXT                                                 NULL,
    created_at  TIMESTAMP                                            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                                            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (content_id, section_no)
);

CREATE TABLE content_logs
(
    id         BIGSERIAL PRIMARY KEY,
    content_id BIGINT REFERENCES contents (id) ON DELETE CASCADE NOT NULL,
    prompt     TEXT                                                 NOT NULL,
    call       JSONB                                                NOT NULL,
    created_at TIMESTAMP                                            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                                            NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE course_photos
(
    id          BIGSERIAL PRIMARY KEY,
    course_id   BIGINT REFERENCES courses (id) ON DELETE CASCADE NOT NULL,
    title       VARCHAR(255)                                         NOT NULL,
    description TEXT                                                 NULL,
    created_at  TIMESTAMP                                            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                                            NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE content_section_photos
(
    id                  BIGSERIAL PRIMARY KEY,
    content_section_id  BIGINT REFERENCES content_sections (id) ON DELETE CASCADE NOT NULL,
    course_photo_id     BIGINT REFERENCES course_photos (id) ON DELETE CASCADE     NOT NULL,
    created_at          TIMESTAMP                                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP                                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (content_section_id, course_photo_id)
);

-- * auto-update function for updated_at timestamps
CREATE OR REPLACE FUNCTION auto_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- * triggers to automatically update updated_at
CREATE TRIGGER auto_updated_at_users
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_topics
    BEFORE UPDATE
    ON topics
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_courses
    BEFORE UPDATE
    ON courses
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_course_managers
    BEFORE UPDATE
    ON course_managers
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_enrolls
    BEFORE UPDATE
    ON enrolls
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_contents
    BEFORE UPDATE
    ON contents
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_content_sections
    BEFORE UPDATE
    ON content_sections
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_content_logs
    BEFORE UPDATE
    ON content_logs
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_course_photos
    BEFORE UPDATE
    ON course_photos
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_content_section_photos
    BEFORE UPDATE
    ON content_section_photos
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE content_section_photos;
DROP TABLE course_photos;
DROP TABLE content_logs;
DROP TABLE content_sections;
DROP TABLE contents;
DROP TABLE enrolls;
DROP TABLE course_managers;
DROP TABLE courses;
DROP TABLE topics;
DROP TABLE users;
DROP FUNCTION auto_updated_at;
-- +goose StatementEnd