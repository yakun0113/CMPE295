import 'package:flutter/material.dart';
import 'package:walking_octopus/components/default_button.dart';
import 'package:walking_octopus/screens/forgetpassword_screen.dart';
import 'package:walking_octopus/screens/loginsuccess_screen.dart';
import 'package:walking_octopus/screens/signup_screen.dart';
import 'package:walking_octopus/size_config.dart';
import 'package:walking_octopus/constants.dart';


class Body extends StatelessWidget {
  
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: SizedBox(
        width: double.infinity,
        child: SingleChildScrollView(
          child: Padding(
            padding: EdgeInsets.symmetric(horizontal: 20),
            child: Column(
              children: [
                Text(
                  "Sign In",
                  style: TextStyle(
                    fontSize: getProportionateScreenWidth(28),
                    fontWeight: FontWeight.bold,
                    color: Colors.black,
                    height: 1.5,
                  ),
                ),
                Text(
                  "You can sign in with social media or with email",
                  textAlign: TextAlign.left,
                ),
                SizedBox(height:40),
                Column(
                  crossAxisAlignment: CrossAxisAlignment.end,
                  children: [
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
                    SizedBox(height: getProportionateScreenHeight(30)),
                    SignForm(),
                    SizedBox(height: getProportionateScreenHeight(30)),
                    NoAccountText(),
                  ],
                ),
              ],
            ),
          ),
        ),
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
            onTap: () => Navigator.pushNamed(context, SignUpScreen.routeName),
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

class SignForm extends StatefulWidget {
  @override
  _SignFormState createState() => _SignFormState();
}

class _SignFormState extends State<SignForm> {
  bool remember = false;
  String email;
  String password;
  final _formKey = GlobalKey<FormState>();
  final List<String> errors = [];
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
                  //suffixIcon: SvgPicture.asset("assets/icons/Mail.svg"),
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
            Container(
              padding: EdgeInsets.symmetric(horizontal: 20, vertical: 5),
              decoration: BoxDecoration(
                  color: Colors.transparent,
                  borderRadius: BorderRadius.all(Radius.circular(80))
              ),
              child: TextFormField(
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
              ),
            ),
            SizedBox(height: getProportionateScreenHeight(20)),
            FormError(errors: errors),
            SizedBox(height: getProportionateScreenHeight(20)),
            DefaultButton(
              text: "Continue",
              press: () {
                if(_formKey.currentState.validate()) {
                  _formKey.currentState.save();
                  Navigator.pushNamed(context, LoginSuccessScreen.routeName);
                }
              },
            ),
            SizedBox(height: getProportionateScreenHeight(20)),
            Column(
              children: [
                Row(
                  children: [
                    Checkbox(
                        value: remember,
                        activeColor: kPrimaryColor,
                        onChanged: (value) {
                          setState(() {
                            remember = value;
                          });
                        }
                    ),
                    Text("Remember me"),
                    Spacer(),
                    GestureDetector(
                        onTap: () => Navigator.pushNamed(context, ForgetPasswordScreen.routeName),
                        child: Text("Forgot Password?", style: Theme.of(context).textTheme.bodyText1,)
                    )
                  ],
                ),
              ],
            )
          ],
        )
    );
  }
}

class FormError extends StatelessWidget {
  const FormError({
    Key key,
    @required this.errors,
  }) : super(key: key);

  final List<String> errors;

  @override
  Widget build(BuildContext context) {
    return Column(
        children: List.generate(
          errors.length, (index) => formErrorText(error: errors[index]),
        )
    );
  }

  Row formErrorText({String error}) {
    return Row(
      children: [
        Image(
            height: 15,
            width: 15,
            image: AssetImage('assets/icons/Error.png')
        ),
        SizedBox(width: 10),
        Text(error),
      ],
    );
  }
}