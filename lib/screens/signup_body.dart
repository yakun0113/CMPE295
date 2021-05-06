import 'package:flutter/material.dart';
import 'package:walking_octopus/components/default_button.dart';
import 'package:walking_octopus/constants.dart';
import 'package:walking_octopus/screens/complete_profile_sceen.dart';
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
              SizedBox(height: SizeConfig.screenHeight * 0.02),
              Text(
                "Register Account",
                style: TextStyle(
                  fontSize: getProportionateScreenWidth(28),
                  fontWeight: FontWeight.bold,
                  color: Colors.black,
                  height: 1.5,
                ),
              ),
              Text(
                "Enter your personal information or continue \nwith social media",
                textAlign: TextAlign.center,
              ),
              SizedBox(height: SizeConfig.screenHeight * 0.08),
              SignUpForm(),
              SizedBox(height: SizeConfig.screenHeight * 0.08),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Image(
                    width: 50,
                    image: AssetImage('assets/icons/google.png'),
                    ),
                  SizedBox(width: 40),
                  Image(
                    width: 50,
                    image: AssetImage('assets/icons/facebook.png')
                  )
                ],
              ),
              SizedBox(height: getProportionateScreenHeight(20)),
              Text(
                "By countinuing your confirm that you agree \nwith our privacy policy",
                textAlign: TextAlign.center,
              )
            ],
          ),
        ),
      ),
    );
  }
}

class SignUpForm extends StatefulWidget {
  @override
  _SignUpFormState createState() => _SignUpFormState();
}

class _SignUpFormState extends State<SignUpForm> {
  final _formKey = GlobalKey<FormState>();
  String email;
  String password;
  String conform_password;
  final List<String> errors = [];
  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        children: [
          buildEmailFormField(),
          SizedBox(height: getProportionateScreenHeight(30)),
          buildPasswordFormField(),
          SizedBox(height: getProportionateScreenHeight(30)),
          buildConfPasswordFormField(),
          FormError(errors: errors),
          SizedBox(height: getProportionateScreenHeight(40)),
          DefaultButton(
            text: "Continue",
            press: (){
              if(_formKey.currentState.validate()) {
                // Go to complete profile page
                Navigator.pushNamed(context, CompleteProfileScreen.routeName);
              }
            }
          )
        ],
      ),
    );
  }

  TextFormField buildConfPasswordFormField() {
    return TextFormField(
              obscureText: true,
              onSaved: (newValue) => conform_password = newValue,
              onChanged: (value) {
                if (password == conform_password) {
                  setState(() {
                    errors.remove(kMatchPassError);
                  });
                } 
                return null;
              },
              validator: (value) {
                if (value.isEmpty && !errors.contains(kPassNullError)) {
                  return "";
                } else if (password != value) {
                  setState(() {
                    errors.add(kMatchPassError);
                  });
                  return "";
                }
                return null;
              },
              decoration: InputDecoration(
                border: InputBorder.none,
                labelText: "Confirm Password",
                floatingLabelBehavior: FloatingLabelBehavior.always,
                hintText: "Re-enter your password",
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

  TextFormField buildPasswordFormField() {
    return TextFormField(
              obscureText: true,
              onSaved: (newValue) => password = newValue,
              onChanged: (value) {
                if (value.isNotEmpty && errors.contains(kPassNullError)) {
                  setState(() {
                    errors.remove(kPassNullError);
                  });
                } else if (value.length >= 8 && errors.contains(kShortPassError)) {
                  setState(() {
                    errors.remove(kShortPassError);
                  });
                }
                password = value;
                return null;
              },
              validator: (value) {
                if (value.isEmpty && !errors.contains(kPassNullError)) {
                  setState(() {
                    errors.add(kPassNullError);
                  });
                  return "";
                } else if (value.length < 8 && !errors.contains(kShortPassError)) {
                  setState(() {
                    errors.add(kShortPassError);
                  });
                  return "";
                }
                return null;
              },
              decoration: InputDecoration(
                border: InputBorder.none,
                labelText: "Password",
                floatingLabelBehavior: FloatingLabelBehavior.always,
                hintText: "Enter your password",
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

  TextFormField buildEmailFormField() {
    return TextFormField(
          keyboardType: TextInputType.emailAddress,
          onSaved: (newValue) => email = newValue,
          onChanged: (value) {
            if (value.isNotEmpty && errors.contains(kEmailNullError)) {
              setState(() {
                errors.remove(kEmailNullError);
              });
            } else if (emailValidatorRegExp.hasMatch(value) && errors.contains(kInvalidEmailError)) {
              setState(() {
                errors.remove(kInvalidEmailError);
              });
            }
            return null;
          },
          validator: (value) {
            if (value.isEmpty && !errors.contains(kEmailNullError)) {
              setState(() {
                errors.add(kEmailNullError);
              });
              return "";
            } else if (!emailValidatorRegExp.hasMatch(value) && !errors.contains(kInvalidEmailError)) {
              setState(() {
                errors.add(kInvalidEmailError);
              });
              return "";
            }
            return null;
          },
          decoration: InputDecoration(
            border: InputBorder.none,
            labelText: "Email",
            floatingLabelBehavior: FloatingLabelBehavior.always,
            hintText: "Enter your email",
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

