import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:smiot/bloc/auth_bloc.dart';
import 'package:smiot/bloc/box_bloc.dart';
import 'package:smiot/bloc/common_bloc.dart';
import 'package:smiot/screens/splash_screen.dart';
import 'package:smiot/widgets/box_card.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({
    Key? key,
  }) : super(key: key);

  static Route<void> route() {
    return MaterialPageRoute<void>(builder: (_) => const HomeScreen());
  }

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  void initState() {
    super.initState();
    context.read<BoxBloc>().add(GetBoxesEvent());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: ListTile(
          trailing: IconButton(
              onPressed: (() {
                context.read<AuthBloc>().add(LogoutEvent());
              }),
              icon: Icon(
                Icons.logout,
                color: Colors.grey[800],
              )),
          title: Text(
            'SMIOT',
            style:
                TextStyle(color: Colors.grey[800], fontWeight: FontWeight.bold),
          ),
        ),
        iconTheme: const IconThemeData(color: Colors.black),
        backgroundColor: Colors.white,
      ),
      body: body(),
    );
  }

  Widget body() => Container(
        padding: const EdgeInsets.only(top: 20, left: 10, right: 10),
        child: BlocBuilder<BoxBloc, MyState>(
          builder: (context, state) {
            if (state is StateLoading) {
              Center(
                child: Container(
                  padding: const EdgeInsets.all(10),
                  child: const CircularProgressIndicator(),
                ),
              );
            } else if (state is GetBoxesStateSuccess) {
              return ListView.builder(
                itemCount: state.boxes.length,
                itemBuilder: ((context, index) => BoxCard(
                      box: state.boxes[index],
                    )),
              );
            }
            return const SizedBox();
          },
        ),
      );
}
