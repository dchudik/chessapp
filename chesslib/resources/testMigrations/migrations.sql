CREATE TABLE IF NOT EXISTS figures_coordinates_errors (
    figure_coordinates varchar(3) NOT NULL,
    king_coordinates varchar(3) NOT NULL,
    error boolean NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS figures_coordinates_under_attack (
    figure_coordinates varchar(3) NOT NULL,
    king_coordinates varchar(3) NOT NULL,
    figure varchar(25) NOT NULL,
    under_attack boolean NOT NULL DEFAULT FALSE
);

