-- Sets up a trigger for the given table to automatically set a column called
-- `updated_at` whenever the row is modified (unless `updated_at` was included
-- in the modified columns)
--
-- # Example
--
-- ```sql
-- CREATE TABLE users (id SERIAL PRIMARY KEY, updated_at TIMESTAMP NOT NULL DEFAULT NOW());
--
-- SELECT diesel_manage_updated_at('users');
-- ```
CREATE OR REPLACE FUNCTION auto_manage_updated_at_and_version(_tbl regclass)
    RETURNS VOID AS
$$
BEGIN
    EXECUTE format('CREATE TRIGGER set_updated_at BEFORE UPDATE ON %s
                      FOR EACH ROW EXECUTE PROCEDURE auto_set_audit_columns()', _tbl);
END;
$$
    LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION auto_set_audit_columns()
    RETURNS trigger AS
$$
BEGIN
    IF (NEW IS DISTINCT FROM OLD)
    THEN
        NEW.updated_at := current_timestamp;
        NEW.version := OLD.version + 1;
    END IF;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;


CREATE TABLE language
(
    id     SERIAL,
    name   TEXT NOT NULL,
    locale TEXT NOT NULL,
    CONSTRAINT pk_language PRIMARY KEY (id),
    CONSTRAINT uk_name UNIQUE (name),
    CONSTRAINT uk_locale UNIQUE (locale)
);


INSERT INTO language (id, name, locale)
VALUES (1, 'English', 'en'),
       (2, 'Italian', 'it'),
       (3, 'German', 'de'),
       (4, 'French', 'fr'),
       (5, 'Portuguese - Brazilian', 'pt_BR'),
       (6, 'Dutch', 'nl'),
       (7, 'Spanish', 'es'),
       (8, 'Norwegian', 'nb_NO'),
       (9, 'Danish', 'da'),
       (10, 'Japanese', 'ja'),
       (11, 'Swedish', 'sv'),
       (12, 'Spanish - Spain', 'es_ES'),
       (13, 'French - Canada', 'fr_CA'),
       (14, 'Lithuanian', 'lt'),
       (15, 'Polish', 'pl'),
       (16, 'Czech', 'cs'),
       (17, 'Croatian', 'hr'),
       (18, 'Albanian', 'sq'),
       (19, 'Greek', 'el'),
       (20, 'English - United Kingdom', 'en_GB'),
       (21, 'Portuguese - Portugal', 'pt_PT'),
       (22, 'Slovenian', 'sl'),
       (23, 'Finnish', 'fi'),
       (24, 'Romanian', 'ro'),
       (25, 'Turkish - Turkey', 'tr_TR'),
       (26, 'Thai', 'th');


CREATE TABLE date_format
(
    id                 SERIAL,
    c_format           TEXT NOT NULL,
    date_picker_format TEXT NOT NULL,
    js_format          TEXT NOT NULL,
    CONSTRAINT pk_date_format PRIMARY KEY (id),
    CONSTRAINT uk_date_format__c_format UNIQUE (c_format)
);


INSERT INTO date_format (id, c_format, date_picker_format, js_format)
VALUES (1, 'd/M/Y', 'dd/M/yyyy', 'DD/MMM/YYYY'),
       (2, 'd-M-Y', 'dd-M-yyyy', 'DD-MMM-YYYY'),
       (3, 'd/F/Y', 'dd/MM/yyyy', 'DD/MMMM/YYYY'),
       (4, 'd-F-Y', 'dd-MM-yyyy', 'DD-MMMM-YYYY'),
       (5, 'M j, Y', 'M d, yyyy', 'MMM D, YYYY'),
       (6, 'F j, Y', 'MM d, yyyy', 'MMMM D, YYYY'),
       (7, 'D M j, Y', 'D MM d, yyyy', 'ddd MMM Do, YYYY'),
       (8, 'Y-m-d', 'yyyy-mm-dd', 'YYYY-MM-DD'),
       (9, 'd-m-Y', 'dd-mm-yyyy', 'DD-MM-YYYY'),
       (10, 'm/d/Y', 'mm/dd/yyyy', 'MM/DD/YYYY'),
       (11, 'd.m.Y', 'dd.mm.yyyy', 'D.MM.YYYY'),
       (12, 'j. M. Y', 'd. M. yyyy', 'DD. MMM. YYYY'),
       (13, 'j. F Y', 'd. MM yyyy', 'DD. MMMM YYYY');


CREATE TABLE datetime_format
(
    id        SERIAL,
    c_format  TEXT NOT NULL,
    js_format TEXT NOT NULL,
    CONSTRAINT pk_datetime_format PRIMARY KEY (id),
    CONSTRAINT uk_datetime_format__c_format UNIQUE (c_format)
);

INSERT INTO datetime_format (id, c_format, js_format)
VALUES (1, 'd/M/Y g:i a', 'DD/MMM/YYYY h:mm:ss a'),
       (2, 'd-M-Y g:i a', 'DD-MMM-YYYY h:mm:ss a'),
       (3, 'd/F/Y g:i a', 'DD/MMMM/YYYY h:mm:ss a'),
       (4, 'd-F-Y g:i a', 'DD-MMMM-YYYY h:mm:ss a'),
       (5, 'M j, Y g:i a', 'MMM D, YYYY h:mm:ss a'),
       (6, 'F j, Y g:i a', 'MMMM D, YYYY h:mm:ss a'),
       (7, 'D M jS, Y g:i a', 'ddd MMM Do, YYYY h:mm:ss a'),
       (8, 'Y-m-d g:i a', 'YYYY-MM-DD h:mm:ss a'),
       (9, 'd-m-Y g:i a', 'DD-MM-YYYY h:mm:ss a'),
       (10, 'm/d/Y g:i a', 'MM/DD/YYYY h:mm:ss a'),
       (11, 'd.m.Y g:i a', 'D.MM.YYYY h:mm:ss a'),
       (12, 'j. M. Y g:i a', 'DD. MMM. YYYY h:mm:ss a'),
       (13, 'j. F Y g:i a', 'DD. MMMM YYYY h:mm:ss a');

CREATE TABLE timezone
(
    id         SERIAL,
    name       TEXT NOT NULL,
    gmt_offset TEXT NOT NULL,
    location   TEXT NOT NULL,
    CONSTRAINT pk_timezone PRIMARY KEY (id),
    CONSTRAINT uk_timezone__code UNIQUE (name)
);

INSERT INTO timezone (id, name, gmt_offset, location)
VALUES (1, 'Pacific/Midway', '-11:00', 'Midway Island'),
       (2, 'US/Samoa', '-11:00', 'Samoa'),
       (3, 'US/Hawaii', '-10:00', 'Hawaii'),
       (4, 'US/Alaska', '-09:00', 'Alaska'),
       (5, 'US/Pacific', '-08:00', 'Pacific Time (US & Canada)'),
       (6, 'America/Tijuana', '-08:00', 'Tijuana'),
       (7, 'US/Arizona', '-07:00', 'Arizona'),
       (8, 'US/Mountain', '-07:00', 'Mountain Time (US & Canada)'),
       (9, 'America/Chihuahua', '-07:00', 'Chihuahua'),
       (10, 'America/Mazatlan', '-07:00', 'Mazatlan'),
       (11, 'America/Mexico_City', '-06:00', 'Mexico City'),
       (12, 'America/Monterrey', '-06:00', 'Monterrey'),
       (13, 'Canada/Saskatchewan', '-06:00', 'Saskatchewan'),
       (14, 'US/Central', '-06:00', 'Central Time (US & Canada)'),
       (15, 'US/Eastern', '-05:00', 'Eastern Time (US & Canada)'),
       (16, 'US/East-Indiana', '-05:00', 'Indiana (East)'),
       (17, 'America/Bogota', '-05:00', 'Bogota'),
       (18, 'America/Lima', '-05:00', 'Lima'),
       (19, 'America/Caracas', '-04:30', 'Caracas'),
       (20, 'Canada/Atlantic', '-04:00', 'Atlantic Time (Canada)'),
       (21, 'America/La_Paz', '-04:00', 'La Paz'),
       (22, 'America/Santiago', '-04:00', 'Santiago'),
       (23, 'Canada/Newfoundland', '-03:30', 'Newfoundland'),
       (24, 'America/Buenos_Aires', '-03:00', 'Buenos Aires'),
       (25, 'America/Godthab', '-03:00', 'Greenland'),
       (26, 'Atlantic/Stanley', '-02:00', 'Stanley'),
       (27, 'Atlantic/Azores', '-01:00', 'Azores'),
       (28, 'Atlantic/Cape_Verde', '-01:00', 'Cape Verde Is.'),
       (29, 'Africa/Casablanca', '00:00', 'Casablanca'),
       (30, 'Europe/Dublin', '00:00', 'Dublin'),
       (31, 'Europe/Lisbon', '00:00', 'Lisbon'),
       (32, 'Europe/London', '00:00', 'London'),
       (33, 'Africa/Monrovia', '00:00', 'Monrovia'),
       (34, 'Europe/Amsterdam', '+01:00', 'Amsterdam'),
       (35, 'Europe/Belgrade', '+01:00', 'Belgrade'),
       (36, 'Europe/Berlin', '+01:00', 'Berlin'),
       (37, 'Europe/Bratislava', '+01:00', 'Bratislava'),
       (38, 'Europe/Brussels', '+01:00', 'Brussels'),
       (39, 'Europe/Budapest', '+01:00', 'Budapest'),
       (40, 'Europe/Copenhagen', '+01:00', 'Copenhagen'),
       (41, 'Europe/Ljubljana', '+01:00', 'Ljubljana'),
       (42, 'Europe/Madrid', '+01:00', 'Madrid'),
       (43, 'Europe/Paris', '+01:00', 'Paris'),
       (44, 'Europe/Prague', '+01:00', 'Prague'),
       (45, 'Europe/Rome', '+01:00', 'Rome'),
       (46, 'Europe/Sarajevo', '+01:00', 'Sarajevo'),
       (47, 'Europe/Skopje', '+01:00', 'Skopje'),
       (48, 'Europe/Stockholm', '+01:00', 'Stockholm'),
       (49, 'Europe/Vienna', '+01:00', 'Vienna'),
       (50, 'Europe/Warsaw', '+01:00', 'Warsaw'),
       (51, 'Europe/Zagreb', '+01:00', 'Zagreb'),
       (52, 'Europe/Athens', '+02:00', 'Athens'),
       (53, 'Europe/Bucharest', '+02:00', 'Bucharest'),
       (54, 'Africa/Cairo', '+02:00', 'Cairo'),
       (55, 'Africa/Harare', '+02:00', 'Harare'),
       (56, 'Europe/Helsinki', '+02:00', 'Helsinki'),
       (57, 'Europe/Istanbul', '+02:00', 'Istanbul'),
       (58, 'Asia/Jerusalem', '+02:00', 'Jerusalem'),
       (59, 'Europe/Kiev', '+02:00', 'Kyiv'),
       (60, 'Europe/Minsk', '+02:00', 'Minsk'),
       (61, 'Europe/Riga', '+02:00', 'Riga'),
       (62, 'Europe/Sofia', '+02:00', 'Sofia'),
       (63, 'Europe/Tallinn', '+02:00', 'Tallinn'),
       (64, 'Europe/Vilnius', '+02:00', 'Vilnius'),
       (65, 'Asia/Baghdad', '+03:00', 'Baghdad'),
       (66, 'Asia/Kuwait', '+03:00', 'Kuwait'),
       (67, 'Africa/Nairobi', '+03:00', 'Nairobi'),
       (68, 'Asia/Riyadh', '+03:00', 'Riyadh'),
       (69, 'Asia/Tehran', '+03:30', 'Tehran'),
       (70, 'Europe/Moscow', '+04:00', 'Moscow'),
       (71, 'Asia/Baku', '+04:00', 'Baku'),
       (72, 'Europe/Volgograd', '+04:00', 'Volgograd'),
       (73, 'Asia/Muscat', '+04:00', 'Muscat'),
       (74, 'Asia/Tbilisi', '+04:00', 'Tbilisi'),
       (75, 'Asia/Yerevan', '+04:00', 'Yerevan'),
       (76, 'Asia/Kabul', '+04:30', 'Kabul'),
       (77, 'Asia/Karachi', '+05:00', 'Karachi'),
       (78, 'Asia/Tashkent', '+05:00', 'Tashkent'),
       (79, 'Asia/Kolkata', '+05:30', 'Kolkata'),
       (80, 'Asia/Kathmandu', '+05:45', 'Kathmandu'),
       (81, 'Asia/Yekaterinburg', '+06:00', 'Ekaterinburg'),
       (82, 'Asia/Almaty', '+06:00', 'Almaty'),
       (83, 'Asia/Dhaka', '+06:00', 'Dhaka'),
       (84, 'Asia/Novosibirsk', '+07:00', 'Novosibirsk'),
       (85, 'Asia/Bangkok', '+07:00', 'Bangkok'),
       (86, 'Asia/Ho_Chi_Minh', '+07:00', 'Ho Chi Minh'),
       (87, 'Asia/Jakarta', '+07:00', 'Jakarta'),
       (88, 'Asia/Krasnoyarsk', '+08:00', 'Krasnoyarsk'),
       (89, 'Asia/Chongqing', '+08:00', 'Chongqing'),
       (90, 'Asia/Hong_Kong', '+08:00', 'Hong Kong'),
       (91, 'Asia/Kuala_Lumpur', '+08:00', 'Kuala Lumpur'),
       (92, 'Australia/Perth', '+08:00', 'Perth'),
       (93, 'Asia/Singapore', '+08:00', 'Singapore'),
       (94, 'Asia/Taipei', '+08:00', 'Taipei'),
       (95, 'Asia/Ulaanbaatar', '+08:00', 'Ulaan Bataar'),
       (96, 'Asia/Urumqi', '+08:00', 'Urumqi'),
       (97, 'Asia/Irkutsk', '+09:00', 'Irkutsk'),
       (98, 'Asia/Seoul', '+09:00', 'Seoul'),
       (99, 'Asia/Tokyo', '+09:00', 'Tokyo'),
       (100, 'Australia/Adelaide', '+09:30', 'Adelaide'),
       (101, 'Australia/Darwin', '+09:30', 'Darwin'),
       (102, 'Asia/Yakutsk', '+10:00', 'Yakutsk'),
       (103, 'Australia/Brisbane', '+10:00', 'Brisbane'),
       (104, 'Australia/Canberra', '+10:00', 'Canberra'),
       (105, 'Pacific/Guam', '+10:00', 'Guam'),
       (106, 'Australia/Hobart', '+10:00', 'Hobart'),
       (107, 'Australia/Melbourne', '+10:00', 'Melbourne'),
       (108, 'Pacific/Port_Moresby', '+10:00', 'Port Moresby'),
       (109, 'Australia/Sydney', '+10:00', 'Sydney'),
       (110, 'Asia/Vladivostok', '+11:00', 'Vladivostok'),
       (111, 'Asia/Magadan', '+12:00', 'Magadan'),
       (112, 'Pacific/Auckland', '+12:00', 'Auckland'),
       (113, 'Pacific/Fiji', '+12:00', 'Fiji');


CREATE TABLE currency
(
    id        SERIAL,
    code      TEXT     NOT NULL,
    symbol    TEXT     NULL,
    name      TEXT     NOT NULL,
    precision SMALLINT NOT NULL,
    format    TEXT     NOT NULL,
    CONSTRAINT pk_currency PRIMARY KEY (id),
    CONSTRAINT uk_currency__code UNIQUE (code)
);


INSERT INTO currency (id, code, symbol, name, precision, format)
VALUES (1, 'USD', '$', 'US Dollar', 2, '###,##.##'),
       (2, 'GBP', '£', 'British Pound', 2, '###,##.##'),
       (3, 'EUR', '€', 'Euro', 2, '###,##.##'),
       (4, 'ZAR', 'R', 'South African Rand', 2, '###,##.##'),
       (5, 'DKK', 'kr', 'Danish Krone', 2, '###,##.##'),
       (6, 'ILS', 'NIS ', 'Israeli Shekel', 2, '###,##.##'),
       (7, 'SEK', 'kr', 'Swedish Krona', 2, '###,##.##'),
       (8, 'KES', 'KSh ', 'Kenyan Shilling', 2, '###,##.##'),
       (9, 'CAD', 'C$', 'Canadian Dollar', 2, '###,##.##'),
       (10, 'PHP', 'P ', 'Philippine Peso', 2, '###,##.##'),
       (11, 'INR', '₹', 'Indian Rupee', 2, '###,##.##'),
       (12, 'AUD', '$', 'Australian Dollar', 2, '###,##.##'),
       (13, 'SGD', '$', 'Singapore Dollar', 2, '###,##.##'),
       (14, 'NOK', 'kr', 'Norske Kroner', 2, '###,##.##'),
       (15, 'NZD', '$', 'New Zealand Dollar', 2, '###,##.##'),
       (16, 'VND', '', 'Vietnamese Dong', 0, '###,##.##'),
       (17, 'CHF', '', 'Swiss Franc', 2, '###,##.##'),
       (18, 'GTQ', 'Q', 'Guatemalan Quetzal', 2, '###,##.##'),
       (19, 'MYR', 'RM', 'Malaysian Ringgit', 2, '###,##.##'),
       (20, 'BRL', 'R$', 'Brazilian Real', 2, '###,##.##'),
       (21, 'THB', '', 'Thai Baht', 2, '###,##.##'),
       (22, 'NGN', '', 'Nigerian Naira', 2, '###,##.##'),
       (23, 'ARS', '$', 'Argentine Peso', 2, '###,##.##'),
       (24, 'BDT', 'Tk', 'Bangladeshi Taka', 2, '###,##.##'),
       (25, 'AED', 'DH ', 'United Arab Emirates Dirham', 2, '###,##.##'),
       (26, 'HKD', '', 'Hong Kong Dollar', 2, '###,##.##'),
       (27, 'IDR', 'Rp', 'Indonesian Rupiah', 2, '###,##.##'),
       (28, 'MXN', '$', 'Mexican Peso', 2, '###,##.##'),
       (29, 'EGP', 'E£', 'Egyptian Pound', 2, '###,##.##'),
       (30, 'COP', '$', 'Colombian Peso', 2, '###,##.##'),
       (31, 'XOF', 'CFA ', 'West African Franc', 2, '###,##.##'),
       (32, 'CNY', 'RMB ', 'Chinese Renminbi', 2, '###,##.##'),
       (33, 'RWF', 'RF ', 'Rwandan Franc', 2, '###,##.##'),
       (34, 'TZS', 'TSh ', 'Tanzanian Shilling', 2, '###,##.##'),
       (35, 'ANG', '', 'Netherlands Antillean Guilder', 2, '###,##.##'),
       (36, 'TTD', 'TT$', 'Trinidad and Tobago Dollar', 2, '###,##.##'),
       (37, 'XCD', 'EC$', 'East Caribbean Dollar', 2, '###,##.##'),
       (38, 'GHS', '', 'Ghanaian Cedi', 2, '###,##.##'),
       (39, 'BGN', '', 'Bulgarian Lev', 2, '###,##.##'),
       (40, 'AWG', 'Afl. ', 'Aruban Florin', 2, '###,##.##'),
       (41, 'TRY', 'TL ', 'Turkish Lira', 2, '###,##.##'),
       (42, 'RON', '', 'Romanian New Leu', 2, '###,##.##'),
       (43, 'HRK', 'kn', 'Croatian Kuna', 2, '###,##.##'),
       (44, 'SAR', '', 'Saudi Riyal', 2, '###,##.##'),
       (45, 'JPY', '¥', 'Japanese Yen', 0, '###,##.##'),
       (46, 'MVR', '', 'Maldivian Rufiyaa', 2, '###,##.##'),
       (47, 'CRC', '', 'Costa Rican Colón', 2, '###,##.##'),
       (48, 'PKR', 'Rs ', 'Pakistani Rupee', 0, '###,##.##'),
       (49, 'PLN', 'zł', 'Polish Zloty', 2, '###,##.##'),
       (50, 'LKR', 'LKR', 'Sri Lankan Rupee', 2, '###,##.##'),
       (51, 'CZK', 'Kč', 'Czech Koruna', 2, '###,##.##'),
       (52, 'UYU', '$', 'Uruguayan Peso', 2, '###,##.##'),
       (53, 'NAD', '$', 'Namibian Dollar', 2, '###,##.##'),
       (54, 'TND', '', 'Tunisian Dinar', 2, '###,##.##'),
       (55, 'RUB', '', 'Russian Ruble', 2, '###,##.##'),
       (56, 'MZN', 'MT', 'Mozambican Metical', 2, '###,##.##'),
       (57, 'OMR', '', 'Omani Rial', 2, '###,##.##'),
       (58, 'UAH', '', 'Ukrainian Hryvnia', 2, '###,##.##'),
       (59, 'MOP', 'MOP$', 'Macanese Pataca', 2, '###,##.##'),
       (60, 'TWD', 'NT$', 'Taiwan New Dollar', 2, '###,##.##'),
       (61, 'DOP', 'RD$', 'Dominican Peso', 2, '###,##.##'),
       (62, 'CLP', '$', 'Chilean Peso', 0, '###,##.##'),
       (63, 'ISK', 'kr', 'Icelandic Króna', 2, '###,##.##'),
       (64, 'PGK', 'K', 'Papua New Guinean Kina', 2, '###,##.##'),
       (65, 'JOD', '', 'Jordanian Dinar', 2, '###,##.##'),
       (66, 'MMK', 'K', 'Myanmar Kyat', 2, '###,##.##'),
       (67, 'PEN', 'S/ ', 'Peruvian Sol', 2, '###,##.##'),
       (68, 'BWP', 'P', 'Botswana Pula', 2, '###,##.##'),
       (69, 'HUF', 'Ft', 'Hungarian Forint', 0, '###,##.##'),
       (70, 'UGX', 'USh ', 'Ugandan Shilling', 2, '###,##.##'),
       (71, 'BBD', '$', 'Barbadian Dollar', 2, '###,##.##'),
       (72, 'BND', 'B$', 'Brunei Dollar', 2, '###,##.##'),
       (73, 'GEL', '', 'Georgian Lari', 2, '###,##.##'),
       (74, 'QAR', 'QR', 'Qatari Riyal', 2, '###,##.##'),
       (75, 'HNL', 'L', 'Honduran Lempira', 2, '###,##.##'),
       (76, 'AFN', '؋', 'Afgani', 2, '###,##.##');

CREATE TABLE country
(
    id        SERIAL,
    code      TEXT     NOT NULL,
    name      TEXT     NOT NULL,
    dial_code SMALLINT NOT NULL,
    currency  TEXT     NOT NULL,
    CONSTRAINT pk_country PRIMARY KEY (id),
    CONSTRAINT uk_country__name UNIQUE (name),
    CONSTRAINT uk_country__code UNIQUE (code)
);

INSERT INTO country (id, code, name, dial_code, currency)
VALUES (1, 'AF', 'Afghanistan', 93, 'AFN'),
       (2, 'AL', 'Albania', 355, 'ALL'),
       (3, 'DZ', 'Algeria', 213, 'DZD'),
       (4, 'AS', 'American Samoa', 1684, 'XXX'),
       (5, 'AD', 'Andorra', 376, 'EUR'),
       (6, 'AO', 'Angola', 244, 'AOA'),
       (7, 'AI', 'Anguilla', 1264, 'XCD'),
       (8, 'AQ', 'Antarctica', 0, 'XXX'),
       (9, 'AG', 'Antigua And Barbuda', 1268, 'XCD'),
       (10, 'AR', 'Argentina', 54, 'ARS'),
       (11, 'AM', 'Armenia', 374, 'AMD'),
       (12, 'AW', 'Aruba', 297, 'AWG'),
       (13, 'AU', 'Australia', 61, 'AUD'),
       (14, 'AT', 'Austria', 43, 'EUR'),
       (15, 'AZ', 'Azerbaijan', 994, 'AZN'),
       (16, 'BS', 'Bahamas The', 1242, 'XXX'),
       (17, 'BH', 'Bahrain', 973, 'BHD'),
       (18, 'BD', 'Bangladesh', 880, 'BDT'),
       (19, 'BB', 'Barbados', 1246, 'BBD'),
       (20, 'BY', 'Belarus', 375, 'BYR'),
       (21, 'BE', 'Belgium', 32, 'EUR'),
       (22, 'BZ', 'Belize', 501, 'BZD'),
       (23, 'BJ', 'Benin', 229, 'XOF'),
       (24, 'BM', 'Bermuda', 1441, 'BMD'),
       (25, 'BT', 'Bhutan', 975, 'BTN'),
       (26, 'BO', 'Bolivia', 591, 'BOB'),
       (27, 'BA', 'Bosnia and Herzegovina', 387, 'BAM'),
       (28, 'BW', 'Botswana', 267, 'BWP'),
       (29, 'BV', 'Bouvet Island', 0, 'XXX'),
       (30, 'BR', 'Brazil', 55, 'BRL'),
       (31, 'IO', 'British Indian Ocean Territory', 246, 'USD'),
       (32, 'BN', 'Brunei', 673, 'BND'),
       (33, 'BG', 'Bulgaria', 359, 'BGN'),
       (34, 'BF', 'Burkina Faso', 226, 'XOF'),
       (35, 'BI', 'Burundi', 257, 'BIF'),
       (36, 'KH', 'Cambodia', 855, 'KHR'),
       (37, 'CM', 'Cameroon', 237, 'XAF'),
       (38, 'CA', 'Canada', 1, 'CAD'),
       (39, 'CV', 'Cape Verde', 238, 'CVE'),
       (40, 'KY', 'Cayman Islands', 1345, 'KYD'),
       (41, 'CF', 'Central African Republic', 236, 'XAF'),
       (42, 'TD', 'Chad', 235, 'XAF'),
       (43, 'CL', 'Chile', 56, 'CLP'),
       (44, 'CN', 'China', 86, 'CNY'),
       (45, 'CX', 'Christmas Island', 61, 'XXX'),
       (46, 'CC', 'Cocos (Keeling) Islands', 672, 'AUD'),
       (47, 'CO', 'Colombia', 57, 'COP'),
       (48, 'KM', 'Comoros', 269, 'KMF'),
       (49, 'CG', 'Congo', 242, 'XXX'),
       (50, 'CD', 'Congo The Democratic Republic Of The', 242, 'XXX'),
       (51, 'CK', 'Cook Islands', 682, 'NZD'),
       (52, 'CR', 'Costa Rica', 506, 'CRC'),
       (53, 'CI', 'Cote D''Ivoire (Ivory Coast)', 225, 'XXX'),
       (54, 'HR', 'Croatia (Hrvatska)', 385, 'XXX'),
       (55, 'CU', 'Cuba', 53, 'CUC'),
       (56, 'CY', 'Cyprus', 357, 'EUR'),
       (57, 'CZ', 'Czech Republic', 420, 'CZK'),
       (58, 'DK', 'Denmark', 45, 'DKK'),
       (59, 'DJ', 'Djibouti', 253, 'DJF'),
       (60, 'DM', 'Dominica', 1767, 'XCD'),
       (61, 'DO', 'Dominican Republic', 1809, 'DOP'),
       (62, 'TP', 'East Timor', 670, 'USD'),
       (63, 'EC', 'Ecuador', 593, 'USD'),
       (64, 'EG', 'Egypt', 20, 'EGP'),
       (65, 'SV', 'El Salvador', 503, 'USD'),
       (66, 'GQ', 'Equatorial Guinea', 240, 'XAF'),
       (67, 'ER', 'Eritrea', 291, 'ERN'),
       (68, 'EE', 'Estonia', 372, 'EUR'),
       (69, 'ET', 'Ethiopia', 251, 'ETB'),
       (70, 'XA', 'External Territories of Australia', 61, 'XXX'),
       (71, 'FK', 'Falkland Islands', 500, 'FKP'),
       (72, 'FO', 'Faroe Islands', 298, 'DKK'),
       (73, 'FJ', 'Fiji Islands', 679, 'XXX'),
       (74, 'FI', 'Finland', 358, 'EUR'),
       (75, 'FR', 'France', 33, 'EUR'),
       (76, 'GF', 'French Guiana', 594, 'XXX'),
       (77, 'PF', 'French Polynesia', 689, 'XPF'),
       (78, 'TF', 'French Southern Territories', 0, 'XXX'),
       (79, 'GA', 'Gabon', 241, 'XAF'),
       (80, 'GM', 'Gambia The', 220, 'XXX'),
       (81, 'GE', 'Georgia', 995, 'GEL'),
       (82, 'DE', 'Germany', 49, 'EUR'),
       (83, 'GH', 'Ghana', 233, 'GHS'),
       (84, 'GI', 'Gibraltar', 350, 'GIP'),
       (85, 'GR', 'Greece', 30, 'EUR'),
       (86, 'GL', 'Greenland', 299, 'XXX'),
       (87, 'GD', 'Grenada', 1473, 'XCD'),
       (88, 'GP', 'Guadeloupe', 590, 'XXX'),
       (89, 'GU', 'Guam', 1671, 'XXX'),
       (90, 'GT', 'Guatemala', 502, 'GTQ'),
       (91, 'XU', 'Guernsey and Alderney', 44, 'XXX'),
       (92, 'GN', 'Guinea', 224, 'GNF'),
       (93, 'GW', 'Guinea-Bissau', 245, 'XOF'),
       (94, 'GY', 'Guyana', 592, 'GYD'),
       (95, 'HT', 'Haiti', 509, 'HTG'),
       (96, 'HM', 'Heard and McDonald Islands', 0, 'XXX'),
       (97, 'HN', 'Honduras', 504, 'HNL'),
       (98, 'HK', 'Hong Kong S.A.R.', 852, 'XXX'),
       (99, 'HU', 'Hungary', 36, 'HUF'),
       (100, 'IS', 'Iceland', 354, 'ISK'),
       (101, 'IN', 'India', 91, 'INR'),
       (102, 'ID', 'Indonesia', 62, 'IDR'),
       (103, 'IR', 'Iran', 98, 'IRR'),
       (104, 'IQ', 'Iraq', 964, 'IQD'),
       (105, 'IE', 'Ireland', 353, 'EUR'),
       (106, 'IL', 'Israel', 972, 'ILS'),
       (107, 'IT', 'Italy', 39, 'EUR'),
       (108, 'JM', 'Jamaica', 1876, 'JMD'),
       (109, 'JP', 'Japan', 81, 'JPY'),
       (110, 'XJ', 'Jersey', 44, 'GBP'),
       (111, 'JO', 'Jordan', 962, 'JOD'),
       (112, 'KZ', 'Kazakhstan', 7, 'KZT'),
       (113, 'KE', 'Kenya', 254, 'KES'),
       (114, 'KI', 'Kiribati', 686, 'AUD'),
       (115, 'KP', 'Korea North', 850, 'XXX'),
       (116, 'KR', 'Korea South', 82, 'XXX'),
       (117, 'KW', 'Kuwait', 965, 'KWD'),
       (118, 'KG', 'Kyrgyzstan', 996, 'KGS'),
       (119, 'LA', 'Laos', 856, 'LAK'),
       (120, 'LV', 'Latvia', 371, 'EUR'),
       (121, 'LB', 'Lebanon', 961, 'LBP'),
       (122, 'LS', 'Lesotho', 266, 'LSL'),
       (123, 'LR', 'Liberia', 231, 'LRD'),
       (124, 'LY', 'Libya', 218, 'LYD'),
       (125, 'LI', 'Liechtenstein', 423, 'CHF'),
       (126, 'LT', 'Lithuania', 370, 'EUR'),
       (127, 'LU', 'Luxembourg', 352, 'EUR'),
       (128, 'MO', 'Macau S.A.R.', 853, 'XXX'),
       (129, 'MK', 'Macedonia', 389, 'XXX'),
       (130, 'MG', 'Madagascar', 261, 'MGA'),
       (131, 'MW', 'Malawi', 265, 'MWK'),
       (132, 'MY', 'Malaysia', 60, 'MYR'),
       (133, 'MV', 'Maldives', 960, 'MVR'),
       (134, 'ML', 'Mali', 223, 'XOF'),
       (135, 'MT', 'Malta', 356, 'EUR'),
       (136, 'XM', 'Man (Isle of)', 44, 'XXX'),
       (137, 'MH', 'Marshall Islands', 692, 'USD'),
       (138, 'MQ', 'Martinique', 596, 'XXX'),
       (139, 'MR', 'Mauritania', 222, 'MRO'),
       (140, 'MU', 'Mauritius', 230, 'MUR'),
       (141, 'YT', 'Mayotte', 269, 'XXX'),
       (142, 'MX', 'Mexico', 52, 'MXN'),
       (143, 'FM', 'Micronesia', 691, 'XXX'),
       (144, 'MD', 'Moldova', 373, 'MDL'),
       (145, 'MC', 'Monaco', 377, 'EUR'),
       (146, 'MN', 'Mongolia', 976, 'MNT'),
       (147, 'MS', 'Montserrat', 1664, 'XCD'),
       (148, 'MA', 'Morocco', 212, 'MAD'),
       (149, 'MZ', 'Mozambique', 258, 'MZN'),
       (150, 'MM', 'Myanmar', 95, 'MMK'),
       (151, 'NA', 'Namibia', 264, 'NAD'),
       (152, 'NR', 'Nauru', 674, 'AUD'),
       (153, 'NP', 'Nepal', 977, 'NPR'),
       (154, 'AN', 'Netherlands Antilles', 599, 'XXX'),
       (155, 'NL', 'Netherlands The', 31, 'XXX'),
       (156, 'NC', 'New Caledonia', 687, 'XPF'),
       (157, 'NZ', 'New Zealand', 64, 'NZD'),
       (158, 'NI', 'Nicaragua', 505, 'NIO'),
       (159, 'NE', 'Niger', 227, 'XOF'),
       (160, 'NG', 'Nigeria', 234, 'NGN'),
       (161, 'NU', 'Niue', 683, 'NZD'),
       (162, 'NF', 'Norfolk Island', 672, 'XXX'),
       (163, 'MP', 'Northern Mariana Islands', 1670, 'XXX'),
       (164, 'NO', 'Norway', 47, 'NOK'),
       (165, 'OM', 'Oman', 968, 'OMR'),
       (166, 'PK', 'Pakistan', 92, 'PKR'),
       (167, 'PW', 'Palau', 680, 'XXX'),
       (168, 'PS', 'Palestinian Territory Occupied', 970, 'XXX'),
       (169, 'PA', 'Panama', 507, 'PAB'),
       (170, 'PG', 'Papua new Guinea', 675, 'PGK'),
       (171, 'PY', 'Paraguay', 595, 'PYG'),
       (172, 'PE', 'Peru', 51, 'PEN'),
       (173, 'PH', 'Philippines', 63, 'PHP'),
       (174, 'PN', 'Pitcairn Island', 0, 'XXX'),
       (175, 'PL', 'Poland', 48, 'PLN'),
       (176, 'PT', 'Portugal', 351, 'EUR'),
       (177, 'PR', 'Puerto Rico', 1787, 'XXX'),
       (178, 'QA', 'Qatar', 974, 'QAR'),
       (179, 'RE', 'Reunion', 262, 'XXX'),
       (180, 'RO', 'Romania', 40, 'RON'),
       (181, 'RU', 'Russia', 70, 'RUB'),
       (182, 'RW', 'Rwanda', 250, 'RWF'),
       (183, 'SH', 'Saint Helena', 290, 'SHP'),
       (184, 'KN', 'Saint Kitts And Nevis', 1869, 'XCD'),
       (185, 'LC', 'Saint Lucia', 1758, 'XCD'),
       (186, 'PM', 'Saint Pierre and Miquelon', 508, 'XXX'),
       (187, 'VC', 'Saint Vincent And The Grenadines', 1784, 'XCD'),
       (188, 'WS', 'Samoa', 684, 'WST'),
       (189, 'SM', 'San Marino', 378, 'EUR'),
       (190, 'ST', 'Sao Tome and Principe', 239, 'STD'),
       (191, 'SA', 'Saudi Arabia', 966, 'SAR'),
       (192, 'SN', 'Senegal', 221, 'XOF'),
       (193, 'RS', 'Serbia', 381, 'RSD'),
       (194, 'SC', 'Seychelles', 248, 'SCR'),
       (195, 'SL', 'Sierra Leone', 232, 'SLL'),
       (196, 'SG', 'Singapore', 65, 'BND'),
       (197, 'SK', 'Slovakia', 421, 'EUR'),
       (198, 'SI', 'Slovenia', 386, 'EUR'),
       (199, 'XG', 'Smaller Territories of the UK', 44, 'XXX'),
       (200, 'SB', 'Solomon Islands', 677, 'SBD'),
       (201, 'SO', 'Somalia', 252, 'SOS'),
       (202, 'ZA', 'South Africa', 27, 'ZAR'),
       (203, 'GS', 'South Georgia', 0, 'XXX'),
       (204, 'SS', 'South Sudan', 211, 'SSP'),
       (205, 'ES', 'Spain', 34, 'EUR'),
       (206, 'LK', 'Sri Lanka', 94, 'LKR'),
       (207, 'SD', 'Sudan', 249, 'SDG'),
       (208, 'SR', 'Suriname', 597, 'SRD'),
       (209, 'SJ', 'Svalbard And Jan Mayen Islands', 47, 'XXX'),
       (210, 'SZ', 'Swaziland', 268, 'SZL'),
       (211, 'SE', 'Sweden', 46, 'SEK'),
       (212, 'CH', 'Switzerland', 41, 'CHF'),
       (213, 'SY', 'Syria', 963, 'SYP'),
       (214, 'TW', 'Taiwan', 886, 'TWD'),
       (215, 'TJ', 'Tajikistan', 992, 'TJS'),
       (216, 'TZ', 'Tanzania', 255, 'TZS'),
       (217, 'TH', 'Thailand', 66, 'THB'),
       (218, 'TG', 'Togo', 228, 'XOF'),
       (219, 'TK', 'Tokelau', 690, 'XXX'),
       (220, 'TO', 'Tonga', 676, 'TOP'),
       (221, 'TT', 'Trinidad And Tobago', 1868, 'TTD'),
       (222, 'TN', 'Tunisia', 216, 'TND'),
       (223, 'TR', 'Turkey', 90, 'TRY'),
       (224, 'TM', 'Turkmenistan', 7370, 'TMT'),
       (225, 'TC', 'Turks And Caicos Islands', 1649, 'USD'),
       (226, 'TV', 'Tuvalu', 688, 'AUD'),
       (227, 'UG', 'Uganda', 256, 'UGX'),
       (228, 'UA', 'Ukraine', 380, 'UAH'),
       (229, 'AE', 'United Arab Emirates', 971, 'AED'),
       (230, 'GB', 'United Kingdom', 44, 'GBP'),
       (231, 'US', 'United States', 1, 'USD'),
       (232, 'UM', 'United States Minor Outlying Islands', 1, 'XXX'),
       (233, 'UY', 'Uruguay', 598, 'UYU'),
       (234, 'UZ', 'Uzbekistan', 998, 'UZS'),
       (235, 'VU', 'Vanuatu', 678, 'VUV'),
       (236, 'VA', 'Vatican City State (Holy See)', 39, 'XXX'),
       (237, 'VE', 'Venezuela', 58, 'VEF'),
       (238, 'VN', 'Vietnam', 84, 'VND'),
       (239, 'VG', 'Virgin Islands (British)', 1284, 'XXX'),
       (240, 'VI', 'Virgin Islands (US)', 1340, 'XXX'),
       (241, 'WF', 'Wallis And Futuna Islands', 681, 'XXX'),
       (242, 'EH', 'Western Sahara', 212, 'XXX'),
       (243, 'YE', 'Yemen', 967, 'YER'),
       (244, 'YU', 'Yugoslavia', 38, 'XXX'),
       (245, 'ZM', 'Zambia', 260, 'ZMW'),
       (246, 'ZW', 'Zimbabwe', 263, 'BWP');


CREATE TABLE user_account
(
    id           BIGSERIAL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by   TEXT                     NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by   TEXT                     NOT NULL,
    version      INTEGER                  NOT NULL DEFAULT 1,
    first_name   TEXT                     NOT NULL,
    last_name    TEXT                     NOT NULL,
    email        TEXT                     NOT NULL,
    phone_number TEXT                     NULL,
    account_type INTEGER                  NULL,
    active       BOOLEAN                  NOT NULL DEFAULT FALSE,
    expires_at   TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT pk_user_account PRIMARY KEY (id)
);

CREATE UNIQUE INDEX uk_user_account__email ON user_account (lower(email));
CREATE UNIQUE INDEX uk_user_account__phone_number ON user_account (lower(phone_number));

SELECT auto_manage_updated_at_and_version('user_account');


CREATE TABLE user_credential
(
    id                        BIGINT,
    updated_at                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    version                   INTEGER                  NOT NULL DEFAULT 1,
    password_hash             TEXT,
    expires_at                TIMESTAMP WITH TIME ZONE,
    invalid_attempts          INT                      NOT NULL DEFAULT 0,
    locked                    BOOLEAN                  NOT NULL DEFAULT FALSE,
    activation_key            TEXT,
    activation_key_expires_at TIMESTAMP WITH TIME ZONE,
    activated                 BOOLEAN                  NOT NULL DEFAULT FALSE,
    reset_key                 TEXT,
    reset_key_expires_at      TIMESTAMP WITH TIME ZONE,
    reset_at                  TIMESTAMP WITH TIME ZONE,
    CONSTRAINT pk_user_credential PRIMARY KEY (id),
    CONSTRAINT fk_user_credential_01 FOREIGN KEY (id) REFERENCES user_account (id)
);

SELECT auto_manage_updated_at_and_version('user_credential');

CREATE TABLE auth_token
(
    id         BIGSERIAL                NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    user_id    BIGINT                   NOT NULL REFERENCES user_account (id),
    token      TEXT                     NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    mobile     BOOLEAN                  NOT NULL DEFAULT FALSE,
    identifier TEXT,
    CONSTRAINT pk_auth_token PRIMARY KEY (id)
);

CREATE TABLE role
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by  TEXT                     NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  TEXT                     NOT NULL,
    version     INTEGER                  NOT NULL DEFAULT 1,
    name        TEXT                     NOT NULL,
    description TEXT                     NOT NULL
);
CREATE UNIQUE INDEX role_uk_01 ON role (lower(name));
SELECT auto_manage_updated_at_and_version('role');


CREATE TABLE permission
(
    id          SERIAL PRIMARY KEY,
    resource    TEXT NOT NULL,
    authority   TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE UNIQUE INDEX permission_uk_01 ON permission (lower(resource), lower(authority));


CREATE TABLE role_permission
(
    role_id       INTEGER NOT NULL,
    permission_id INTEGER NOT NULL,
    CONSTRAINT role_permissions_pk PRIMARY KEY (role_id, permission_id),
    CONSTRAINT role_permission_fk_01 FOREIGN KEY (role_id) REFERENCES role (id),
    CONSTRAINT role_permission_fk_02 FOREIGN KEY (permission_id) REFERENCES permission (id)
);


CREATE TABLE user_role
(
    user_id BIGINT  NOT NULL,
    role_id INTEGER NOT NULL,
    CONSTRAINT user_role_pk PRIMARY KEY (user_id, role_id),
    CONSTRAINT user_role_fk_01 FOREIGN KEY (role_id) REFERENCES role (id),
    CONSTRAINT user_role_fk_02 FOREIGN KEY (user_id) REFERENCES user_account (id)
);

CREATE TABLE user_group
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by  TEXT                     NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  TEXT                     NOT NULL,
    version     INTEGER                  NOT NULL DEFAULT 1,
    name        TEXT                     NOT NULL,
    description TEXT                     NOT NULL
);

CREATE UNIQUE INDEX user_group_uk ON user_group (lower(name));

CREATE TABLE user_group_user
(
    group_id INTEGER NOT NULL,
    user_id  BIGINT  NOT NULL,
    CONSTRAINT user_group_user_pk PRIMARY KEY (group_id, user_id),
    CONSTRAINT user_group_user_fk_01 FOREIGN KEY (group_id) REFERENCES user_group (id),
    CONSTRAINT user_group_user_fk_02 FOREIGN KEY (user_id) REFERENCES user_account (id)
);

CREATE TABLE user_group_role
(
    group_id INTEGER NOT NULL,
    role_id  INTEGER NOT NULL,
    CONSTRAINT user_group_role_pk PRIMARY KEY (group_id, role_id),
    CONSTRAINT user_group_role_fk_01 FOREIGN KEY (group_id) REFERENCES user_group (id),
    CONSTRAINT user_group_role_fk_02 FOREIGN KEY (role_id) REFERENCES role (id)
);


