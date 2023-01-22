import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smiot/bloc/auth_bloc.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/screens/home_screen.dart';
import 'package:smiot/screens/login_screen.dart';
import 'package:smiot/widgets/splash.dart';

class AuthGuard extends StatefulWidget {
  final Widget child;
  const AuthGuard({
    Key? key,
    required this.child,
  }) : super(key: key);

  @override
  State<AuthGuard> createState() => _AuthGuardState();
}

class _AuthGuardState extends State<AuthGuard> {
  @override
  void initState() {
    super.initState();
    context.read<AuthBloc>().add(GetTokenEvent());
  }

  final _navigatorKey = GlobalKey<NavigatorState>();

  NavigatorState get _navigator => _navigatorKey.currentState!;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      navigatorKey: _navigatorKey,
      title: 'MATEE KLOUD',
      builder: (context, child) {
        return BlocListener<AuthBloc, MyState>(
          listener: ((context, state) {
            if (state is AuthState) {
              if (state.token.isEmpty) {
                _navigator.pop();
                _navigator.push(LoginScreen.route());
                return;
              } else {
                _navigator.pop();
                _navigator.push(HomeScreen.route());
                return;
              }
            }
          }),
          child: child,
        );
      },
      onGenerateRoute: ((_) => LoginScreen.route()),
    );
  }
}
