from flask import Flask, request, jsonify
from collections import Counter
import requests
from bs4 import BeautifulSoup
from datetime import datetime

app = Flask(__name__)

# Constants
URL = 'https://www.fakenamegenerator.com/'
HEADERS = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"
}
BLOOD_TYPES = ["O-", "O+", "A-", "A+", "B-", "B+", "AB-", "AB+"]
BLOOD_TYPE_COMPATIBILITY = {
    "O-": ["O-", "O+", "A-", "A+", "B-", "B+", "AB-", "AB+"],
    "O+": ["O+", "A+", "B+", "AB+"],
    "A-": ["A-", "A+", "AB-", "AB+"],
    "A+": ["A+", "AB+"],
    "B-": ["B-", "B+", "AB-", "AB+"],
    "B+": ["B+", "AB+"],
    "AB-": ["AB-", "AB+"],
    "AB+": ["AB+"]
}


def get_zodiac_sign(month, day):
    if (month == 1 and day >= 20) or (month == 2 and day <= 18):
        return 'Aquarius'
    elif (month == 2 and day >= 19) or (month == 3 and day <= 20):
        return 'Pisces'
    elif (month == 3 and day >= 21) or (month == 4 and day <= 19):
        return 'Aries'
    elif (month == 4 and day >= 20) or (month == 5 and day <= 20):
        return 'Taurus'
    elif (month == 5 and day >= 21) or (month == 6 and day <= 20):
        return 'Gemini'
    elif (month == 6 and day >= 21) or (month == 7 and day <= 22):
        return 'Cancer'
    elif (month == 7 and day >= 23) or (month == 8 and day <= 22):
        return 'Leo'
    elif (month == 8 and day >= 23) or (month == 9 and day <= 22):
        return 'Virgo'
    elif (month == 9 and day >= 23) or (month == 10 and day <= 22):
        return 'Libra'
    elif (month == 10 and day >= 23) or (month == 11 and day <= 21):
        return 'Scorpio'
    elif (month == 11 and day >= 22) or (month == 12 and day <= 21):
        return 'Sagittarius'
    else:
        return 'Capricorn'


# Helper function to calculate age
def calculate_age(born):
    today = datetime.today()
    return today.year - born.year - ((today.month, today.day) < (born.month, born.day))


def fetch_page_content(url, headers):
    response = requests.get(url, headers=headers)
    if response.status_code != 200:
        print(
            f"Request to {url} failed with status code {response.status_code}")
        return None
    return BeautifulSoup(response.content, 'html.parser')


def get_fake_name():
    soup = fetch_page_content(URL, HEADERS)
    if soup:
        name_element = soup.find('div', class_='info')
        if name_element:
            address_element = name_element.find('div', class_='address')
            if address_element:
                name = address_element.find('h3').text.strip()
                first_name, last_name = name.split(' ')[0], name.split(' ')[-1]
                return first_name, last_name
    return None, None


def get_old_person():
    MAX_TRIES = 10
    tries = 0

    while tries < MAX_TRIES:
        soup = fetch_page_content(URL, HEADERS)
        tries += 1
        if soup:
            name = soup.find('div', class_='address').find(
                'h3').get_text(strip=True)
            dob_element = soup.find('dt', string='Birthday')
            if dob_element:
                dob = dob_element.find_next_sibling('dd').get_text(strip=True)
                dob_date = datetime.strptime(dob, "%B %d, %Y")
                formatted_dob = dob_date.strftime("%d/%m/%Y")
                dob_year = dob_date.year

                if dob_year >= 1940 and dob_year <= 1950:
                    first_name, middle_name, last_name = name.split()
                    return {
                        "first_name": first_name,
                        "middle_name": middle_name,
                        "last_name": last_name,
                        "date_of_birth": formatted_dob,
                        "tropical_zodiac": get_zodiac_sign(dob_date.month, dob_date.day)
                    }
    print(f"No suitable person found after {MAX_TRIES} tries.")
    return None


@app.route('/old_person', methods=['GET'])
def old_person():
    person = get_old_person()
    return jsonify(person), 200


@app.route('/top_used_words', methods=['GET'])
def top_used_words():
    number_of_names = request.args.get('number_of_names', default=5, type=int)
    words = []
    for _ in range(number_of_names):
        first_name, last_name = get_fake_name()
        words.extend([first_name, last_name])
    counter = Counter(words)
    most_common_words = counter.most_common(10)
    return {word: count for word, count in most_common_words}


def get_blood_donors(blood_type):
    donors = []
    attempts = 0
    while len(donors) < 20 and attempts < 100:
        person = get_old_person()
        if not person:
            attempts += 1
            continue
        person_blood_type = person["Blood Type"]
        person_age = datetime.now().year - \
            datetime.strptime(person["Date Of Birth"], "%d/%m/%Y").year
        if person_age >= 18 and person_age <= 45 and blood_type in BLOOD_TYPE_COMPATIBILITY[person_blood_type]:
            donors.append(person)
        attempts += 1
    return donors


@app.route('/get_blood_donors', methods=['GET'])
def get_blood_donors_route():
    required_blood_type = request.args.get('blood_type', type=str)
    if required_blood_type not in BLOOD_TYPES:
        return jsonify({'error': 'Invalid blood type provided'}), 400
    suitable_donors = []
    for i in range(10):
        response = fetch_page_content(URL, HEADERS)
        if response:
            name = response.find('div', class_='address').find(
                'h3').get_text(strip=True)
            dob_element = response.find('dt', string='Birthday')
            if dob_element:
                dob = dob_element.find_next_sibling('dd').get_text(strip=True)
                dob_date = datetime.strptime(dob, "%B %d, %Y")
                blood_type_element = response.find('dt', string='Blood type')
                if blood_type_element:
                    blood_type = blood_type_element.find_next_sibling(
                        'dd').get_text(strip=True)
                    if blood_type in BLOOD_TYPE_COMPATIBILITY[required_blood_type]:
                        first_name, middle_name, last_name = name.split()
                        suitable_donors.append({
                            "first_name": first_name,
                            "middle_name": middle_name,
                            "last_name": last_name,
                            "date_of_birth": dob_date.strftime("%d/%m/%Y"),
                            "blood_type": blood_type,
                            "age": calculate_age(dob_date)
                        })
    return jsonify(suitable_donors), 200


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=80)
