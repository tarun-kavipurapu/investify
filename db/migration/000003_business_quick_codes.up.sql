CREATE TABLE
    "bk_quick_codes" (
        "quick_code_id" BIGSERIAL PRIMARY KEY,
        "type" VARCHAR(10) NOT NULL,
        "name" VARCHAR(255) NOT NULL,
        "code" VARCHAR(50) NOT NULL,
        "description" TEXT,
        "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        UNIQUE ("code")
    );

-- State Codes
INSERT INTO
    bk_quick_codes (type, name, code, description)
VALUES
    ('STATE', 'Andhra Pradesh', 'AP', 'State in India'),
    (
        'STATE',
        'Arunachal Pradesh',
        'AR',
        'State in India'
    ),
    ('STATE', 'Assam', 'AS', 'State in India'),
    ('STATE', 'Bihar', 'BR', 'State in India'),
    ('STATE', 'Chhattisgarh', 'CG', 'State in India'),
    ('STATE', 'Goa', 'GA', 'State in India'),
    ('STATE', 'Gujarat', 'GJ', 'State in India'),
    ('STATE', 'Haryana', 'HR', 'State in India'),
    (
        'STATE',
        'Himachal Pradesh',
        'HP',
        'State in India'
    ),
    ('STATE', 'Jharkhand', 'JH', 'State in India'),
    ('STATE', 'Karnataka', 'KA', 'State in India'),
    ('STATE', 'Kerala', 'KL', 'State in India'),
    ('STATE', 'Madhya Pradesh', 'MP', 'State in India'),
    ('STATE', 'Maharashtra', 'MH', 'State in India'),
    ('STATE', 'Manipur', 'MN', 'State in India'),
    ('STATE', 'Meghalaya', 'ML', 'State in India'),
    ('STATE', 'Mizoram', 'MZ', 'State in India'),
    ('STATE', 'Nagaland', 'NL', 'State in India'),
    ('STATE', 'Odisha', 'OR', 'State in India'),
    ('STATE', 'Punjab', 'PB', 'State in India'),
    ('STATE', 'Rajasthan', 'RJ', 'State in India'),
    ('STATE', 'Sikkim', 'SK', 'State in India'),
    ('STATE', 'Tamil Nadu', 'TN', 'State in India'),
    ('STATE', 'Telangana', 'TS', 'State in India'),
    ('STATE', 'Tripura', 'TR', 'State in India'),
    ('STATE', 'Uttar Pradesh', 'UP', 'State in India'),
    ('STATE', 'Uttarakhand', 'UK', 'State in India'),
    ('STATE', 'West Bengal', 'WB', 'State in India');

-- Business Domains
INSERT INTO
    bk_quick_codes (type, name, code, description)
VALUES
    (
        'DOMAIN',
        'Restaurant',
        'REST',
        'Restaurant Business Domain'
    ),
    (
        'DOMAIN',
        'Retail',
        'RETL',
        'Retail Business Domain'
    ),
    (
        'DOMAIN',
        'Information Technology',
        'IT',
        'Information Technology Business Domain'
    ),
    (
        'DOMAIN',
        'Healthcare',
        'HLTH',
        'Healthcare Business Domain'
    ),
    (
        'DOMAIN',
        'Education',
        'EDUC',
        'Education Business Domain'
    ),
    (
        'DOMAIN',
        'Manufacturing',
        'MNFG',
        'Manufacturing Business Domain'
    ),
    (
        'DOMAIN',
        'Finance',
        'FIN',
        'Finance Business Domain'
    ),
    (
        'DOMAIN',
        'Real Estate',
        'RE',
        'Real Estate Business Domain'
    ),
    (
        'DOMAIN',
        'Transportation',
        'TRANS',
        'Transportation Business Domain'
    ),
    (
        'DOMAIN',
        'Tourism',
        'TOUR',
        'Tourism Business Domain'
    );