import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smiot/bloc/auth_bloc.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/screens/home_screen.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({
    Key? key,
  }) : super(key: key);

  static Route<void> route() {
    return MaterialPageRoute<void>(builder: (_) => const LoginScreen());
  }

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

final key = GlobalKey<FormState>();
var email = '';
var password = '';

class _LoginScreenState extends State<LoginScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Login',
          style:
              TextStyle(color: Colors.grey[800], fontWeight: FontWeight.bold),
        ),
        iconTheme: const IconThemeData(color: Colors.black),
        backgroundColor: Colors.white,
      ),
      body: SingleChildScrollView(
        child: Container(
          padding: const EdgeInsets.all(10),
          child: Column(
            children: [
              buildForm(),
              ElevatedButton(
                onPressed: (() {
                  if (!(key.currentState?.validate() ?? false)) {
                    return;
                  }

                  key.currentState?.save();
                  context
                      .read<AuthBloc>()
                      .add(LoginEvet(email: email, password: password));
                }),
                child: BlocBuilder<AuthBloc, MyState>(
                  builder: ((context, state) => state is StateLoading
                      ? Text(state.toString())
                      : const Text('เข้าสู่ระบบ')),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget buildForm() => Form(
      key: key,
      child: Column(
        children: [
          TextFormField(
            decoration: const InputDecoration(labelText: 'email'),
            maxLength: 50,
            keyboardType: TextInputType.emailAddress,
            onSaved: (value) => email = value ??= '',
            validator: (value) {
              value ??= '';
              if (value.isEmpty) return 'กรุณากรอกอีเมล';
              if (!RegExp(r'^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$')
                  .hasMatch(value)) {
                return 'อีเมลไม่ถูกต้อง';
              }

              return null;
            },
          ),
          TextFormField(
            decoration: const InputDecoration(labelText: 'password'),
            keyboardType: TextInputType.visiblePassword,
            onSaved: (value) => password = value ??= '',
            validator: (value) {
              value ??= '';
              if (value.isEmpty) return 'กรุณากรอกรหัสผ่าน';
              return null;
            },
          )
        ],
      ));
}
