import 'package:flutter/material.dart';
import 'package:walking_octopus/components/default_button.dart';
import 'package:walking_octopus/screens/forgetpassword_screen.dart';
import 'package:walking_octopus/screens/signin_body.dart';
import 'package:walking_octopus/constants.dart';
import 'package:walking_octopus/size_config.dart';


class Body extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal:20),
          child: Column(
            children: [
              SizedBox(height: 40),
              Text(
                "Forget Password",
                style: TextStyle(
                  fontSize: getProportionateScreenWidth(28),
                  fontWeight: FontWeight.bold,
                  color: Colors.black,
                  height: 1.5,
                ),
              ),
              Text(
                "Enter your email and we will send you a link",
                textAlign: TextAlign.center,
              ),
              SizedBox(height: 80),
              ForgetPassForm(),
            ],
          ),
        ),
      ),
    );
  }
}

class ForgetPassForm extends StatefulWidget {
  @override
  _ForgetPassFormState createState() => _ForgetPassFormState();
}

class _ForgetPassFormState extends State<ForgetPassForm> {
  final _formKey = GlobalKey<FormState>();
  List<String> errors = [];
  String email;
  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      child: Column(
        children: [
          Container(
            padding: EdgeInsets.symmetric(horizontal: 20, vertical: 5),
            decoration: BoxDecoration(
                color: Colors.transparent,
                borderRadius: BorderRadius.all(Radius.circular(80))
            ),
            child: TextFormField(
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
                } else if (!emailValidatorRegExp.hasMatch(value) && !errors.contains(kInvalidEmailError)) {
                  setState(() {
                    errors.add(kInvalidEmailError);
                  });
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
            ),
          ),
          SizedBox(height: 20),
          FormError(errors: errors),
          SizedBox(height: 80),
          DefaultButton(
            text: "Continue",
            press: () {
              if (_formKey.currentState.validate()){
                // Do what you want to do
              }
            },
          ),
          SizedBox(height: 80),
          NoAccountText(),
        ],
      ),
    );
  }
}
class NoAccountText extends StatelessWidget {
  const NoAccountText({
    Key key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            "Have no account?",
            style: TextStyle(
              fontSize: getProportionateScreenWidth(18),
            ),
          ),
          GestureDetector(
            onTap: () => Navigator.popAndPushNamed(context, ForgetPasswordScreen.routeName),
            child: Text(
                "Create account",
                style: TextStyle(
                  fontSize: getProportionateScreenWidth(18),
                  color: kPrimaryColor,
                )
            ),
          )
        ]
    );
  }
}