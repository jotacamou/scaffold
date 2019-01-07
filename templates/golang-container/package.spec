Name:           {{.Program}}
Version:        %{_version}
Release:        1
Summary:        Binary RPM package for {{.Name}}

License:        GPLv3
Source0:        ../SOURCES/%{name}

%description
Some long description.

%install
rm -rf %{buildroot}
install -d -m 0755 %{buildroot}/usr/local/bin
install -m 0755 %{SOURCE0} %{buildroot}/usr/local/bin

%clean
rm -rf %{buildroot}

%files
/usr/local/bin/%{name}

%changelog
