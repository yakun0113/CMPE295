import 'package:flutter/material.dart';
import 'package:walking_octopus/components/default_button.dart';
import 'package:walking_octopus/constants.dart';
import 'package:walking_octopus/screens/otp_screen.dart';
import 'package:walking_octopus/screens/signin_body.dart';
import 'package:walking_octopus/size_config.dart';

class Body extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: Padding(
        padding: EdgeInsets.symmetric(horizontal: getProportionateScreenWidth(20)),
        child: SingleChildScrollView(
          child: Column(
            children: [
              Text(
                "Complete Profile",
                style: TextStyle(
                  fontSize: getProportionateScreenWidth(28),
                  fontWeight: FontWeight.bold,
                  color: Colors.black,
                  height: 1.5,
                ),
              ),
              Text(
                "Complete your details or continue \nwith social media",
                textAlign: TextAlign.center,
              ),
              SizedBox(height: SizeConfig.screenHeight * 0.05), // 5%
              CompleteProfileForm(),
              SizedBox(height: getProportionateScreenHeight(30)),
              Text(
                "By continuing your confirm that you agree \nwith our private policy",
                textAlign: TextAlign.center,
              )
            ],
          ),
        ),
      ),
    );
  }
}

class CompleteProfileForm extends StatefulWidget {
  @override
  _CompleteProfileFormState createState() => _CompleteProfileFormState();
}

class _CompleteProfileFormState extends State<CompleteProfileForm> {
  final _formKey = GlobalKey<FormState>();
  final List<String> errors = [];
  String firstName;
  String lastName;
  String phoneNumber;
  String address;

  void addError({String error}) {
    if (!errors.contains(error))
      setState(() {
        errors.add(error);
      });
  }

  void removeError({String error}) {
    if (errors.contains(error))
      setState(() {
        errors.remove(error);
      });
  }


  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        children: [
          buildFirstNameFormField(),
          SizedBox(height: getProportionateScreenHeight(30)),
          buildLastNameFormField(),
          SizedBox(height: getProportionateScreenHeight(30)),
          buildPhoneNumberFormField(),
          SizedBox(height: getProportionateScreenHeight(30)),
          buildAddressFormField(),
          FormError(errors: errors),
          SizedBox(height: getProportionateScreenHeight(30)),
          DefaultButton(
            text: "Continue",
            press: () {
              if (_formKey.currentState.validate()) {
                // Go to OTP screen
                Navigator.pushNamed(context, OtpScreen.routeName);
              }
            }
          )
        ],
      )
    );
  }

  TextFormField buildAddressFormField() {
    return TextFormField(
          onSaved: (newValue) => address = newValue,
          onChanged: (value) {
            if (value.isNotEmpty) {
              setState(() {
                errors.remove(kAddressNullError);
              });
            }
            return null;
          },
          validator: (value) {
            if (value.isEmpty) {
              setState(() {
                errors.add(kAddressNullError);
              });
              return "";
            }
            return null;
          },
          decoration: InputDecoration(
            border: InputBorder.none,
            labelText: "Address",
            floatingLabelBehavior: FloatingLabelBehavior.always,
            hintText: "Enter your Address",
            contentPadding: EdgeInsets.symmetric(horizontal: 36, vertical: 18),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
          ),
        );
  }

  TextFormField buildPhoneNumberFormField() {
    return TextFormField(
          keyboardType: TextInputType.number,
          onSaved: (newValue) => phoneNumber = newValue,
          onChanged: (value) {
            if (value.isNotEmpty) {
              setState(() {
                errors.remove(kPhoneNumberNullError);
              });
            }
            return null;
          },
          validator: (value) {
            if (value.isEmpty) {
              setState(() {
                errors.add(kPhoneNumberNullError);
              });
              return "";
            }
            return null;
          },
          decoration: InputDecoration(
            border: InputBorder.none,
            labelText: "Phone Number",
            floatingLabelBehavior: FloatingLabelBehavior.always,
            hintText: "Enter your phone number",
            contentPadding: EdgeInsets.symmetric(horizontal: 36, vertical: 18),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
          ),
        );
  }

  TextFormField buildLastNameFormField() {
    return TextFormField(
          onSaved: (newValue) => lastName = newValue,
          decoration: InputDecoration(
            border: InputBorder.none,
            labelText: "Last Name",
            floatingLabelBehavior: FloatingLabelBehavior.always,
            hintText: "Enter your last name",
            contentPadding: EdgeInsets.symmetric(horizontal: 36, vertical: 18),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
          ),
        );
  }

  TextFormField buildFirstNameFormField() {
    return TextFormField(
          onSaved: (newValue) => firstName = newValue,
          onChanged: (value) {
            if (value.isNotEmpty) {
              setState(() {
                errors.remove(kNameNullError);
              });
            }
            return null;
          },
          validator: (value) {
            if (value.isEmpty) {
              setState(() {
                errors.add(kNameNullError);
              });
              return "";
            }
            return null;
          },
          decoration: InputDecoration(
            border: InputBorder.none,
            labelText: "First Name",
            floatingLabelBehavior: FloatingLabelBehavior.always,
            hintText: "Enter your first name",
            contentPadding: EdgeInsets.symmetric(horizontal: 36, vertical: 18),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(26),
              borderSide: BorderSide(color: kTextColor),
              gapPadding: 8,
            ),
          ),
        );
  }
}

















