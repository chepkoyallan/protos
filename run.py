from proto.personal_pb2 import Personal
from proto import personal_pb2

allan = Personal()
allan.first_name = "Allan"
allan.second_name = "Kiplang'at"
allan.linkedin_profile = "https://www.linkedin.com/in/allankiplangat/"
allan.header_finished = True
phone_numbers = allan.phone_numbers
phone_numbers.append("0718150768")
allan.education_level = personal_pb2.BACHELORS

allan.single_address.address_one = "Kaminoge"
allan.single_address.address_two = "Takatsu"

#repeated
address_list = allan.addresses.add(address_one = "Kaminoge", address_two = "Takatsu")

